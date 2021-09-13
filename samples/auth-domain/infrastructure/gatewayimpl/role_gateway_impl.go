package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/role"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"golang.org/x/net/context"
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
