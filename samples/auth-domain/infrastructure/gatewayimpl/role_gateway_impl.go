package gatewayimpl

import (
	"context"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/domain/role"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/convertor"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/repo"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
)

func init() {
	ioc.AppendInjection(func(reRepo *repo.UserRoleRepo) role.RoleGateway {
		return &RoleGatewayImpl{
			reRepo: reRepo,
		}
	})
}

type RoleGatewayImpl struct {
	reRepo *repo.UserRoleRepo
}

func (r *RoleGatewayImpl) FindByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]*role.Role, error) {
	pos, err := r.reRepo.ListByUserId(dto.TenantId, userId)
	if err != nil {
		return nil, err
	}
	return convertor.PosToBatchRoleEntitys(pos)
}
