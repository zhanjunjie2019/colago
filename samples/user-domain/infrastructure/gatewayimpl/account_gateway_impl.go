package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/model/entity"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo"
)

func init() {
	_ = ioc.InjectSimpleBean(new(AccountGatewayImpl))
}

type AccountGatewayImpl struct {
	accRepo *repo.AccountRepo `ij:"repo.AccountRepo"`
}

func (a *AccountGatewayImpl) AccRepo() *repo.AccountRepo {
	return a.accRepo
}

func (a *AccountGatewayImpl) SetAccRepo(accRepo *repo.AccountRepo) {
	a.accRepo = accRepo
}

func (a AccountGatewayImpl) New() ioc.AbsBean {
	return a
}

func (a AccountGatewayImpl) FindAccountByAccKey(dto *client.DTO, accKey string) (*entity.Account, error) {
	account, err := a.accRepo.FindByAccKey(dto.TenantId, accKey)
	if err != nil {
		return nil, err
	}
	return convertor.ToAccountEntity(account)
}
