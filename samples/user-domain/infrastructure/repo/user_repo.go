package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func init() {
	_ = ioc.InjectSimpleBean(new(UserRepo))
}

type UserRepo struct {
	postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (u *UserRepo) Postgres() *postgres.Postgres {
	return u.postgres
}

func (u *UserRepo) SetPostgres(postgres *postgres.Postgres) {
	u.postgres = postgres
}

func (u *UserRepo) New() ioc.AbsBean {
	return u
}

func (u *UserRepo) FindById(tenantId uint64, userId uint64) (*po.UserInfo, error) {
	user := &po.UserInfo{
		TenantId: tenantId,
	}
	err := u.postgres.FindOne(user, "id=? and deleted=?", userId, 0)
	return user, err
}
