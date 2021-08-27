package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo/po"
)

func init() {
	_ = ioc.InjectSimpleBean(new(UserRoleRepo))
}

type UserRoleRepo struct {
	postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (u *UserRoleRepo) Postgres() *postgres.Postgres {
	return u.postgres
}

func (u *UserRoleRepo) SetPostgres(postgres *postgres.Postgres) {
	u.postgres = postgres
}

func (u *UserRoleRepo) New() ioc.AbsBean {
	return u
}

func (u *UserRoleRepo) ListByUserId(tenantId uint64, userId uint64) ([]po.RelationUserRole, error) {
	r := &po.RelationUserRole{
		TenantId: tenantId,
	}
	roles := make([]po.RelationUserRole, 0)
	err := u.postgres.FindList(r, roles, "user_id=? and deleted=?", userId, 0)
	return roles, err
}
