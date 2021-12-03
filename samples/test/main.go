package main

import (
	"fmt"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/common/protoactor"
	"github.com/zhanjunjie2019/colago/common/sentinel"
	"github.com/zhanjunjie2019/colago/samples/test/testcase"
	"golang.org/x/net/context"
)

func init() {
	ioc.BatchProvideFinal()
}

func main() {
	protoactor.InitConsulActorClient(
		"127.0.0.1:8500",
		"colago-samples",
		0,
	)

	err := ioc.GetContainer().Invoke(func(
		sentFilter *sentinel.SentinelFilter) {
		protoactor.InitClientFilters(
			sentFilter,
		)
	})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	tenantid := uint64(7)

	// 用户服务创建新的租户
	testcase.InitUserTenant(context.Background(), tenantid)
	// 权限服务创建新的租户
	testcase.InitAuthTenant(context.Background(), tenantid)
	// 创建新的用户
	testcase.CreateUserAction(context.Background(), tenantid)
	// 用户登录行为
	testcase.LoginAction(context.Background(), tenantid)
}
