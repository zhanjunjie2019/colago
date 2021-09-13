package convertor

import (
	"github.com/zhanjunjie2019/colago/samples/auth-domain/domain/role"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/repo/po"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
)

func CodesToBatchRoleEntitys(codes []string) ([]*role.Role, error) {
	roleEntitys := make([]*role.Role, 0)
	for _, c := range codes {
		roleEntity := new(role.Role)
		roleEntity.SetCode(c)
		roleEntitys = append(roleEntitys, roleEntity)
	}
	return roleEntitys, nil
}

func PosToBatchRoleEntitys(roles []*po.RelationUserRole) ([]*role.Role, error) {
	roleEntitys := make([]*role.Role, 0)
	for _, r := range roles {
		entity, err := PoToRoleEntity(r)
		if err != nil {
			return nil, err
		}
		roleEntitys = append(roleEntitys, entity)
	}
	return roleEntitys, nil
}

func PoToRoleEntity(r *po.RelationUserRole) (*role.Role, error) {
	roleEntity := new(role.Role)
	roleEntity.SetCode(r.RoleCode)
	return roleEntity, nil
}

func EntitysToBatchRolePos(dto *client.DTO, userid uint64, roles []*role.Role) ([]*po.RelationUserRole, error) {
	pos := make([]*po.RelationUserRole, 0)
	for _, r := range roles {
		rolePo, err := EntityToRolePo(dto, userid, r)
		if err != nil {
			return nil, err
		}
		pos = append(pos, rolePo)
	}
	return pos, nil
}

func EntityToRolePo(dto *client.DTO, userid uint64, role *role.Role) (*po.RelationUserRole, error) {
	return &po.RelationUserRole{
		UserId:   userid,
		RoleCode: role.Code(),
		TenantId: dto.TenantId,
	}, nil
}
