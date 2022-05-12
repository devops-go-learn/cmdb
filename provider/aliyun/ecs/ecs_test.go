package ecs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/cmdb/provider/aliyun/ecs"
)

var (
	operater *ecs.EcsOperator
)

func TestQuery(t *testing.T) {
	req := ecs.NewPageQueryRequest()
	pager := operater.PageQuery(req)
	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribe(t *testing.T) {
	ins, err := operater.Describe(&ecs.DescribeRequest{Id: "i-bp1f6d1sbq8s9mm59jeu"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operater = aliyun.O().EcsOperator()
}
