package testcase

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	userclient "e.coding.net/double-j/ego/colago/samples/user-client"
	"fmt"
	"time"
)

func CreateUserAction(tenantid uint64) {
	bean, err := ioc.GetBean("userclient.UserClient")
	if err != nil {
		fmt.Println("创建新的用户:" + err.Error())
		panic(err)
	}
	usercli := bean.(*userclient.UserClient)
	sns := time.Now()
	err = usercli.CreateUserAction(&client.CreateUserCmd{
		Dto: &client.DTO{
			TenantId: tenantid,
		},
		AccType:     1,
		AccKey:      "test_user",
		Password:    "123456",
		FirstName:   "姓氏",
		LastName:    "名",
		Age:         18,
		BirthdayTs:  1234567890123,
		Email:       "123456@dianchu.com",
		PhoneNumber: "12345678912",
		Roles:       []string{"ADMIN", "USER"},
		Auths:       []string{"READ", "WRITE"},
	})
	if err != nil {
		fmt.Println("创建新的用户:" + err.Error())
		panic(err)
	}
	fmt.Println("创建新的用户耗时：" + time.Since(sns).String())
}
