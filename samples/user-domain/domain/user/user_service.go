package user

import (
	"fmt"
	"github.com/zhanjunjie2019/colago/common/codec"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/common/model"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"github.com/zhanjunjie2019/colago/samples/user-domain/domain/account"
	"golang.org/x/net/context"
	"strings"
	"time"
)

func init() {
	ioc.AppendInjection(func(
		accountGateway account.AccountGateway,
		userGateway UserGateway) *UserService {
		return &UserService{
			accountGateway: accountGateway,
			userGateway:    userGateway,
		}
	})
}

type UserService struct {
	accountGateway account.AccountGateway
	userGateway    UserGateway
}

func (u *UserService) LoginAction(ctx context.Context, dto *client.DTO, accKey string, pwd string) (*model.TokenData, error) {
	acc, err := u.accountGateway.FindAccountByAccKey(ctx, dto, accKey)
	if err != nil {
		return nil, err
	}
	if !acc.Enable() {
		return nil, fmt.Errorf("账号状态异常")
	}
	if strings.EqualFold(codec.ToSHA1(pwd), acc.Password()) {
		user, err := u.userGateway.FindByAccount(ctx, dto, acc)
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
