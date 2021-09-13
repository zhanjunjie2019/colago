package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo"
	"golang.org/x/net/context"
)

func init() {
	ioc.AppendInjection(func(accRepo *repo.AccountRepo) account.AccountGateway {
		return &AccountGatewayImpl{
			accRepo: accRepo,
		}
	})
}

type AccountGatewayImpl struct {
	accRepo *repo.AccountRepo
}

func (a *AccountGatewayImpl) FindAccountByAccKey(ctx context.Context, dto *client.DTO, accKey string) (*account.Account, error) {
	acc, err := a.accRepo.FindByAccKey(dto.TenantId, accKey)
	if err != nil {
		return nil, err
	}
	return convertor.PoToAccountEntity(ctx, acc)
}
