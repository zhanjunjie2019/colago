package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/user"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(UserGatewayImpl))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type UserGatewayImpl struct {
	roleRepo *repo.UserRoleRepo `ij:"repo.UserRoleRepo"`
	authRepo *repo.UserAuthRepo `ij:"repo.UserAuthRepo"`
}

func (u *UserGatewayImpl) RoleRepo() *repo.UserRoleRepo {
	return u.roleRepo
}

func (u *UserGatewayImpl) SetRoleRepo(roleRepo *repo.UserRoleRepo) {
	u.roleRepo = roleRepo
}

func (u *UserGatewayImpl) AuthRepo() *repo.UserAuthRepo {
	return u.authRepo
}

func (u *UserGatewayImpl) SetAuthRepo(authRepo *repo.UserAuthRepo) {
	u.authRepo = authRepo
}

func (u *UserGatewayImpl) New() ioc.AbsBean {
	return u
}

func (u *UserGatewayImpl) SaveRoleAuth(dto *client.DTO, user *user.User) error {
	rolePos, err := convertor.EntitysToBatchRolePos(dto, user.Id(), user.Roles())
	if err != nil {
		return err
	}
	_, err = u.roleRepo.InsertBatch(rolePos)
	if err != nil {
		return err
	}
	authPos, err := convertor.EntitysToBatchAuthPos(dto, user.Id(), user.Auths())
	if err != nil {
		return err
	}
	_, err = u.authRepo.InsertBatch(authPos)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserGatewayImpl) FindById(dto *client.DTO, userId uint64) (*user.User, error) {
	userBean, err := domain.GetDomainFactory().Create("user.User")
	if err != nil {
		return nil, err
	}
	usr := userBean.(*user.User)
	usr.SetId(userId)
	usr.SetDto(dto)
	return usr, nil
}
