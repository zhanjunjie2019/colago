package testcase

import (
	"fmt"
	"github.com/zhanjunjie2019/colago/common/ioc"
	authclient "github.com/zhanjunjie2019/colago/samples/auth-client"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"golang.org/x/net/context"
	"time"
)

func InitAuthTenant(ctx context.Context, tenantid uint64) {
	var authcli *authclient.AuthClient
	err := ioc.GetContainer().Invoke(func(a *authclient.AuthClient) {
		authcli = a
	})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

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
