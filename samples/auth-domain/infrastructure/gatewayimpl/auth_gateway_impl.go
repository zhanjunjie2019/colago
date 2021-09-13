package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/auth"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"golang.org/x/net/context"
)

func init() {
	ioc.AppendInjection(func(reRepo *repo.UserAuthRepo) auth.AuthGateway {
		return &AuthGatewayImpl{
			reRepo: reRepo,
		}
	})
}

type AuthGatewayImpl struct {
	reRepo *repo.UserAuthRepo
}

func (a *AuthGatewayImpl) FindByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]*auth.Auth, error) {
	pos, err := a.reRepo.ListByUserId(dto.TenantId, userId)
	if err != nil {
		return nil, err
	}
	return convertor.PosToBatchAuthEntitys(pos)
}
