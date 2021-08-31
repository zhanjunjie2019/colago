package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/auth"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(AuthGatewayImpl))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type AuthGatewayImpl struct {
	reRepo *repo.UserAuthRepo `ij:"repo.UserAuthRepo"`
}

func (a *AuthGatewayImpl) ReRepo() *repo.UserAuthRepo {
	return a.reRepo
}

func (a *AuthGatewayImpl) SetReRepo(reRepo *repo.UserAuthRepo) {
	a.reRepo = reRepo
}

func (a *AuthGatewayImpl) New() ioc.AbsBean {
	return a
}

func (a *AuthGatewayImpl) FindByUserId(dto *client.DTO, userId uint64) ([]*auth.Auth, error) {
	pos, err := a.reRepo.ListByUserId(dto.TenantId, userId)
	if err != nil {
		return nil, err
	}
	return convertor.PosToBatchAuthEntitys(pos)
}
