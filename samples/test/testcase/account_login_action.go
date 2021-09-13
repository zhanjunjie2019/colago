package testcase

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	userclient "e.coding.net/double-j/ego/colago/samples/user-client"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func LoginAction(ctx context.Context, tenantid uint64) {
	var usercli *userclient.UserClient
	err := ioc.GetContainer().Invoke(func(u *userclient.UserClient) {
		usercli = u
	})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	sns := time.Now()
	action, err := usercli.LoginAction(ctx, &client.UserLoginCmd{
		Dto: &client.DTO{
			TenantId: tenantid,
		},
		AccKey:   "test_user",
		Password: "123456",
	})
	if err != nil {
		fmt.Println("用户登录行为:" + err.Error())
		panic(err)
	}
	marshal, _ := json.Marshal(action)
	fmt.Println(string(marshal))
	fmt.Println("用户登录行为耗时：" + time.Since(sns).String())
}
