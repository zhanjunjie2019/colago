package gatewayimpl

import (
	"context"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/domain/auth"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/convertor"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/repo"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
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
