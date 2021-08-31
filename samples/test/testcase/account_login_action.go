package testcase

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	userclient "e.coding.net/double-j/ego/colago/samples/user-client"
	"encoding/json"
	"fmt"
)

func LoginAction(tenantid uint64) {
	action, err := userclient.LoginAction(&client.UserLoginCmd{
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
}
