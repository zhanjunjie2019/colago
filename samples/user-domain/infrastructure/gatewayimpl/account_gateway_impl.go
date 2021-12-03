package gatewayimpl

import (
	"context"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"github.com/zhanjunjie2019/colago/samples/user-domain/domain/account"
	"github.com/zhanjunjie2019/colago/samples/user-domain/infrastructure/convertor"
	"github.com/zhanjunjie2019/colago/samples/user-domain/infrastructure/repo"
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
