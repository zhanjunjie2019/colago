package testcase

import (
	authclient "e.coding.net/double-j/ego/colago/samples/auth-client"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
)

func InitAuthTenant(tenantid uint64) {
	err := authclient.InitAuthTenant(&client.AuthTenantInitCmd{
		Dto: &client.DTO{
			TenantId: tenantid,
		},
		TenantId: tenantid,
	})
	if err != nil {
		fmt.Println("权限服务创建新的租户:" + err.Error())
		panic(err)
	}
}
