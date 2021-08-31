package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(AccountGatewayImpl))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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

func (a *AccountGatewayImpl) New() ioc.AbsBean {
	return a
}

func (a *AccountGatewayImpl) FindAccountByAccKey(dto *client.DTO, accKey string) (*account.Account, error) {
	acc, err := a.accRepo.FindByAccKey(dto.TenantId, accKey)
	if err != nil {
		return nil, err
	}
	return convertor.PoToAccountEntity(acc)
}
