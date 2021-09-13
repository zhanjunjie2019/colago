package testcase

import (
	"fmt"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	userclient "github.com/zhanjunjie2019/colago/samples/user-client"
	"golang.org/x/net/context"
	"time"
)

func CreateUserAction(ctx context.Context, tenantid uint64) {
	var usercli *userclient.UserClient
	err := ioc.GetContainer().Invoke(func(u *userclient.UserClient) {
		usercli = u
	})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	sns := time.Now()
	err = usercli.CreateUserAction(ctx, &client.CreateUserCmd{
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
