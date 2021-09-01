package testcase

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	userclient "e.coding.net/double-j/ego/colago/samples/user-client"
	"fmt"
	"strconv"
	"time"
)

func InitUserTenant(tenantid uint64) {
	sns := time.Now().Nanosecond()
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
	ens := time.Now().Nanosecond()
	fmt.Println("用户服务创建新的租户耗时：" + strconv.Itoa((ens-sns)/1000000) + "ms")
}
