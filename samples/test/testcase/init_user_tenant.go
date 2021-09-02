package testcase

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	userclient "e.coding.net/double-j/ego/colago/samples/user-client"
	"fmt"
	"time"
)

func InitUserTenant(tenantid uint64) {
	bean, err := ioc.GetBean("userclient.UserClient")
	if err != nil {
		fmt.Println("用户服务创建新的租户:" + err.Error())
		panic(err)
	}
	usercli := bean.(*userclient.UserClient)
	sns := time.Now()
	err = usercli.InitUserTenant(&client.UserTenantInitCmd{
		Dto: &client.DTO{
			TenantId: tenantid,
		},
		TenantId: tenantid,
	})
	if err != nil {
		fmt.Println("用户服务创建新的租户:" + err.Error())
		panic(err)
	}
	fmt.Println("用户服务创建新的租户耗时：" + time.Since(sns).String())

}
