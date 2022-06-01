package bss_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider/aliyun"
	op "github.com/infraboard/cmdb/provider/aliyun/bss"
)

var (
	operator *op.BssOperator
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	req.Month = "2022-05"
	req.ProductCode = "dysms"

	pager := operator.PageQuery(req)
	for pager.Next() {
		set := bill.NewBillSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()

	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().BssOperator()
}
