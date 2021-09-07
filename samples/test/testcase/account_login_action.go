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
	bean, err := ioc.GetBean("userclient.UserClient")
	if err != nil {
		fmt.Println("创建新的用户:" + err.Error())
		panic(err)
	}
	usercli := bean.(*userclient.UserClient)
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
