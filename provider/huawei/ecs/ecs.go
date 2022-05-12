package ecs

import (
	"strconv"
	"time"

	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperator(client *ecs.EcsClient) *EcsOperator {
	return &EcsOperator{
		client:        client,
		log:           zap.L().Named("Huawei ECS"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type EcsOperator struct {
	client *ecs.EcsClient
	log    logger.Logger
	*resource.AccountGetter
}

func (o *EcsOperator) transferSet(list *[]model.ServerDetail) *host.HostSet {
	set := host.NewHostSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *EcsOperator) transferOne(ins model.ServerDetail) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.Vendor_HUAWEI
	h.Base.Zone = ins.OSEXTAZavailabilityZone
	h.Base.CreateAt = o.parseTime(ins.Created)
	h.Base.Id = ins.Id

	h.Information.Category = ins.Flavor.Name
	h.Information.ExpireAt = o.parseTime(ins.AutoTerminateTime)
	h.Information.Name = ins.Name
	h.Information.Description = utils.PtrStrV(ins.Description)
	h.Information.Status = ins.Status
	h.Information.Tags = o.transferTags(ins.Tags)
	h.Information.PrivateIp, h.Information.PublicIp = o.parseIp(ins.Addresses)
	h.Information.PayType = ins.Metadata["charging_mode"]
	h.Information.SyncAccount = o.GetAccountId()

	h.Describe.SerialNumber = ins.Id
	h.Describe.Cpu, _ = strconv.ParseInt(ins.Flavor.Vcpus, 10, 64)
	h.Describe.Memory, _ = strconv.ParseInt(ins.Flavor.Ram, 10, 64)
	h.Describe.OsType = ins.Metadata["os_type"]
	h.Describe.OsName = ins.Metadata["image_name"]
	h.Describe.ImageId = ins.Image.Id
	h.Describe.KeyPairName = []string{ins.KeyName}
	return h
}

func (o *EcsOperator) transferTags(tags *[]string) (ret []*resource.Tag) {
	if tags == nil {
		return
	}

	t := *tags

	for i := range t {
		ret = append(ret, resource.NewThirdTag("ecs", t[i]))
	}

	return
}

func (o *EcsOperator) parseTime(t string) int64 {
	if t == "" {
		return 0
	}

	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}

func (o *EcsOperator) parseIp(address map[string][]model.ServerAddress) (privateIps []string, publicIps []string) {
	for _, addrs := range address {
		for i := range addrs {
			switch *addrs[i].OSEXTIPStype {
			case model.GetServerAddressOSEXTIPStypeEnum().FIXED:
				privateIps = append(privateIps, addrs[i].Addr)
			case model.GetServerAddressOSEXTIPStypeEnum().FLOATING:
				publicIps = append(publicIps, addrs[i].Addr)
			}
		}
	}
	return
}
