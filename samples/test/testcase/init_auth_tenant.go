package testcase

import (
	authclient "e.coding.net/double-j/ego/colago/samples/auth-client"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"strconv"
	"time"
)

func InitAuthTenant(tenantid uint64) {
	sns := time.Now().Nanosecond()
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
	ens := time.Now().Nanosecond()
	fmt.Println("权限服务创建新的租户耗时：" + strconv.Itoa((ens-sns)/1000000) + "ms")
}
