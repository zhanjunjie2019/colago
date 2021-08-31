package main

import (
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/samples/test/testcase"
)

func main() {
	protoactor.InitConsulActorClient(
		"127.0.0.1:8500",
		"colago-samples",
		0,
	)

	tenantid := uint64(2)

	// 用户服务创建新的租户
	testcase.InitUserTenant(tenantid)
	// 权限服务创建新的租户
	testcase.InitAuthTenant(tenantid)
	// 创建新的用户
	testcase.CreateUserAction(tenantid)
	// 用户登录行为
	testcase.LoginAction(tenantid)
}
