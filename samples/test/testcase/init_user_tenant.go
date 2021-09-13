package testcase

import (
	"fmt"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	userclient "github.com/zhanjunjie2019/colago/samples/user-client"
	"golang.org/x/net/context"
	"time"
)

func InitUserTenant(ctx context.Context, tenantid uint64) {
	var usercli *userclient.UserClient
	err := ioc.GetContainer().Invoke(func(u *userclient.UserClient) {
		usercli = u
	})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	sns := time.Now()
	err = usercli.InitUserTenant(ctx, &client.UserTenantInitCmd{
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
