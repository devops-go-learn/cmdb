package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	cmdbOss "github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewOssOperator(client *oss.Client) *OssOperator {
	return &OssOperator{
		client: client,
		log:    zap.L().Named("ALI Oss"),
	}
}

type OssOperator struct {
	client *oss.Client
	log    logger.Logger
}

func (o *OssOperator) transferSet(items oss.ListBucketsResult) *cmdbOss.BucketSet {
	set := cmdbOss.NewBucketSet()
	for _, item := range items.Buckets {
		set.Add(o.transferBucket(item))
	}
	return set
}

func (o *OssOperator) transferBucket(ins oss.BucketProperties) *cmdbOss.Bucket {
	r := cmdbOss.NewDefaultBucket()

	b := r.Base
	b.Vendor = resource.Vendor_ALIYUN
	b.Region = ins.Location
	b.Id = ins.Name

	return r
}
