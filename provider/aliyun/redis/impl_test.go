package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider/aliyun"
)

var (
	operator provider.RedisOperator
	ctx      = context.Background()
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRedisWithRate(5)
	pager := operator.PageQueryRedis(req)
	for pager.Next() {
		set := redis.NewSet()
		if err := pager.Scan(ctx, set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeRedis(t *testing.T) {
	req := provider.NewDescribeRequest("xx")
	ins, err := operator.DescribeRedis(ctx, req)
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
	operator = aliyun.O().RedisOperator()
}
