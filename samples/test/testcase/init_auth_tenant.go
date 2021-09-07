package testcase

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	authclient "e.coding.net/double-j/ego/colago/samples/auth-client"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func InitAuthTenant(ctx context.Context, tenantid uint64) {
	bean, err := ioc.GetBean("authclient.AuthClient")
	if err != nil {
		fmt.Println("权限服务创建新的租户:" + err.Error())
		panic(err)
	}
	authcli := bean.(*authclient.AuthClient)
	sns := time.Now()
	err = authcli.InitAuthTenant(ctx, &client.AuthTenantInitCmd{
		Dto: &client.DTO{
			TenantId: tenantid,
		},
		TenantId: tenantid,
	})
	if err != nil {
		fmt.Println("权限服务创建新的租户:" + err.Error())
		panic(err)
	}
	fmt.Println("权限服务创建新的租户耗时：" + time.Since(sns).String())
}
