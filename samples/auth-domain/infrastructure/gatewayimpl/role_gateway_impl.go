package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/role"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"golang.org/x/net/context"
)

func init() {
	err := ioc.InjectSimpleBean(new(RoleGatewayImpl))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type RoleGatewayImpl struct {
	ReRepo *repo.UserRoleRepo `ij:"repo.UserRoleRepo"`
}

func (r *RoleGatewayImpl) New() ioc.AbsBean {
	return r
}

func (r *RoleGatewayImpl) FindByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]*role.Role, error) {
	pos, err := r.ReRepo.ListByUserId(dto.TenantId, userId)
	if err != nil {
		return nil, err
	}
	return convertor.PosToBatchRoleEntitys(pos)
}
