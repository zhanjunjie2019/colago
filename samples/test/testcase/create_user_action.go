package testcase

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	userclient "e.coding.net/double-j/ego/colago/samples/user-client"
	"fmt"
)

func CreateUserAction(tenantid uint64) {
	err := userclient.CreateUserAction(&client.CreateUserCmd{
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
}
