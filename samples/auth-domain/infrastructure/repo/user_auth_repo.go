package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo/po"
)

func init() {
	_ = ioc.InjectSimpleBean(new(UserAuthRepo))
}

type UserAuthRepo struct {
	postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (u *UserAuthRepo) Postgres() *postgres.Postgres {
	return u.postgres
}

func (u *UserAuthRepo) SetPostgres(postgres *postgres.Postgres) {
	u.postgres = postgres
}

func (u *UserAuthRepo) New() ioc.AbsBean {
	return u
}

func (u *UserAuthRepo) ListByUserId(tenantId uint64, userId uint64) ([]po.RelationUserAuth, error) {
	r := &po.RelationUserAuth{
		TenantId: tenantId,
	}
	auths := make([]po.RelationUserAuth, 0)
	err := u.postgres.FindList(r, auths, "user_id=? and deleted=?", userId, 0)
	return auths, err
}
