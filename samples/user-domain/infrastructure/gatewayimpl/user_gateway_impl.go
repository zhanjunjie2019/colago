package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/model/entity"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo"
)

func init() {
	_ = ioc.InjectSimpleBean(new(UserGatewayImpl))
}

type UserGatewayImpl struct {
	reRepo   *repo.AccountUserRepo `ij:"repo.AccountUserRepo"`
	userRepo *repo.UserRepo        `ij:"repo.UserRepo"`
}

func (u *UserGatewayImpl) ReRepo() *repo.AccountUserRepo {
	return u.reRepo
}

func (u *UserGatewayImpl) SetReRepo(reRepo *repo.AccountUserRepo) {
	u.reRepo = reRepo
}

func (u *UserGatewayImpl) UserRepo() *repo.UserRepo {
	return u.userRepo
}

func (u *UserGatewayImpl) SetUserRepo(userRepo *repo.UserRepo) {
	u.userRepo = userRepo
}

func (u *UserGatewayImpl) New() ioc.AbsBean {
	return u
}

func (u *UserGatewayImpl) FindByAccount(dto *client.DTO, acc *entity.Account) (*entity.User, error) {
	rela, err := u.reRepo.FindByAccountId(dto.TenantId, acc.Id())
	if err != nil {
		return nil, err
	}
	user, err := u.userRepo.FindById(dto.TenantId, rela.UserId)
	if err != nil {
		return nil, err
	}
	return convertor.ToUserEntity(user)
}
