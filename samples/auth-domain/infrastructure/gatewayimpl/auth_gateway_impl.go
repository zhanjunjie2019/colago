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
	ReRepo *repo.UserAuthRepo `ij:"repo.UserAuthRepo"`
}

func (a *AuthGatewayImpl) New() ioc.AbsBean {
	return a
}

func (a *AuthGatewayImpl) FindByUserId(dto *client.DTO, userId uint64) ([]*auth.Auth, error) {
	pos, err := a.ReRepo.ListByUserId(dto.TenantId, userId)
	if err != nil {
		return nil, err
	}
	return convertor.PosToBatchAuthEntitys(pos)
}
