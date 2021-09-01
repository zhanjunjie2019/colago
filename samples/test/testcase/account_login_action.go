package testcase

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	userclient "e.coding.net/double-j/ego/colago/samples/user-client"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func LoginAction(tenantid uint64) {
	sns := time.Now().Nanosecond()
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
	ens := time.Now().Nanosecond()
	fmt.Println("用户登录行为耗时：" + strconv.Itoa((ens-sns)/1000000) + "ms")
}
