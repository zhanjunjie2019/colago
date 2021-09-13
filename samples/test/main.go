package main

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/common/sentinel"
	"e.coding.net/double-j/ego/colago/common/skywalking"
	"e.coding.net/double-j/ego/colago/samples/test/testcase"
	"fmt"
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

	err := skywalking.NewGlobalTracer("test-main", "127.0.0.1:11800")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = ioc.GetContainer().Invoke(func(
		sentFilter *sentinel.SentinelFilter,
		skyFilter *skywalking.SkyFilter) {
		protoactor.InitClientFilters(
			sentFilter,
			skyFilter,
		)
	})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	span, err := skywalking.NewRootLocalSpan()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer func() {
		span.End(err)
	}()

	tenantid := uint64(7)

	// 用户服务创建新的租户
	testcase.InitUserTenant(span.Ctx(), tenantid)
	// 权限服务创建新的租户
	testcase.InitAuthTenant(span.Ctx(), tenantid)
	// 创建新的用户
	testcase.CreateUserAction(span.Ctx(), tenantid)
	// 用户登录行为
	testcase.LoginAction(span.Ctx(), tenantid)
}
