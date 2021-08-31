package testcase

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	userclient "e.coding.net/double-j/ego/colago/samples/user-client"
	"fmt"
)

func InitUserTenant(tenantid uint64) {
	err := userclient.InitUserTenant(&client.UserTenantInitCmd{
		Dto: &client.DTO{
			TenantId: tenantid,
		},
		TenantId: tenantid,
	})
	if err != nil {
		fmt.Println("用户服务创建新的租户:" + err.Error())
		panic(err)
	}
}
