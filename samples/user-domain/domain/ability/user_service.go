package ability

import (
	"e.coding.net/double-j/ego/colago/common/codec"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/model"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/gateway"
	"fmt"
	"strings"
	"time"
)

func init() {
	_ = ioc.InjectSimpleBean(new(UserService))
}

type UserService struct {
	accountGateway gateway.AccountGateway `ij:"gatewayimpl.AccountGatewayImpl"`
	userGateway    gateway.UserGateway    `ij:"gatewayimpl.UserGatewayImpl"`
}

func (u *UserService) UserGateway() gateway.UserGateway {
	return u.userGateway
}

func (u *UserService) SetUserGateway(userGateway gateway.UserGateway) {
	u.userGateway = userGateway
}

func (u *UserService) AccountGateway() gateway.AccountGateway {
	return u.accountGateway
}

func (u *UserService) SetAccountGateway(accountGateway gateway.AccountGateway) {
	u.accountGateway = accountGateway
}

func (u *UserService) New() ioc.AbsBean {
	return u
}

func (u *UserService) LoginAction(dto *client.DTO, accKey string, pwd string) (*model.TokenData, error) {
	account, err := u.accountGateway.FindAccountByAccKey(dto, accKey)
	if err != nil {
		return nil, err
	}
	if !account.Enable() {
		return nil, fmt.Errorf("账号状态异常")
	}
	if strings.EqualFold(codec.ToSHA1(pwd), account.Password()) {
		user, err := u.userGateway.FindByAccount(dto, account)
		if err != nil {
			return nil, err
		}
		if user.Status() == 1 {
			return nil, fmt.Errorf("用户状态异常")
		}
		return &model.TokenData{
			TenantId: dto.TenantId,
			UserId:   user.Id(),
			UserName: user.FirstName() + "." + user.LastName(),
			HasRoles: user.Roles(),
			HasAuths: user.Auths(),
			Exp:      uint64(time.Now().Unix() + 7200),
		}, nil
	} else {
		return nil, fmt.Errorf("密码不正确")
	}
}
