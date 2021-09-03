package user

import (
	"e.coding.net/double-j/ego/colago/common/codec"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/model"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"fmt"
	"strings"
	"time"
)

func init() {
	err := ioc.InjectSimpleBean(new(UserService))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type UserService struct {
	AccountGateway account.AccountGateway `ij:"gatewayimpl.AccountGatewayImpl"`
	UserGateway    UserGateway            `ij:"gatewayimpl.UserGatewayImpl"`
}

func (u *UserService) New() ioc.AbsBean {
	return u
}

func (u *UserService) LoginAction(dto *client.DTO, accKey string, pwd string) (*model.TokenData, error) {
	acc, err := u.AccountGateway.FindAccountByAccKey(dto, accKey)
	if err != nil {
		return nil, err
	}
	if !acc.Enable() {
		return nil, fmt.Errorf("账号状态异常")
	}
	if strings.EqualFold(codec.ToSHA1(pwd), acc.Password()) {
		user, err := u.UserGateway.FindByAccount(dto, acc)
		if err != nil {
			return nil, err
		}
		if user.Status() != 1 {
			return nil, fmt.Errorf("用户状态异常")
		}
		return &model.TokenData{
			TenantId:  dto.TenantId,
			AccountId: acc.Id(),
			UserId:    user.Id(),
			UserName:  user.FirstName() + "." + user.LastName(),
			HasRoles:  user.Roles(),
			HasAuths:  user.Auths(),
			Exp:       uint64(time.Now().Unix() + 7200),
		}, nil
	} else {
		return nil, fmt.Errorf("密码不正确")
	}
}
