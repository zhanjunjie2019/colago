package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/user"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"golang.org/x/net/context"
)

func init() {
	err := ioc.InjectSimpleBean(new(UserGatewayImpl))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type UserGatewayImpl struct {
	RoleRepo *repo.UserRoleRepo `ij:"repo.UserRoleRepo"`
	AuthRepo *repo.UserAuthRepo `ij:"repo.UserAuthRepo"`
}

func (u *UserGatewayImpl) New() ioc.AbsBean {
	return u
}

func (u *UserGatewayImpl) SaveRoleAuth(ctx context.Context, dto *client.DTO, user *user.User) error {
	rolePos, err := convertor.EntitysToBatchRolePos(dto, user.Id(), user.Roles())
	if err != nil {
		return err
	}
	_, err = u.RoleRepo.InsertBatch(rolePos)
	if err != nil {
		return err
	}
	authPos, err := convertor.EntitysToBatchAuthPos(dto, user.Id(), user.Auths())
	if err != nil {
		return err
	}
	_, err = u.AuthRepo.InsertBatch(authPos)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserGatewayImpl) FindById(ctx context.Context, dto *client.DTO, userId uint64) (*user.User, error) {
	userBean, err := domain.GetDomainFactory().Create("user.User")
	if err != nil {
		return nil, err
	}
	usr := userBean.(*user.User)
	usr.SetId(userId)
	usr.SetDto(dto)
	return usr, nil
}
