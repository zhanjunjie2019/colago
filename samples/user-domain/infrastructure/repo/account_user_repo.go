package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func init() {
	_ = ioc.InjectSimpleBean(new(AccountUserRepo))
}

type AccountUserRepo struct {
	postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (a *AccountUserRepo) Postgres() *postgres.Postgres {
	return a.postgres
}

func (a *AccountUserRepo) SetPostgres(postgres *postgres.Postgres) {
	a.postgres = postgres
}

func (a *AccountUserRepo) New() ioc.AbsBean {
	return a
}

func (a *AccountUserRepo) FindByAccountId(tenantId uint64, accId uint64) (*po.RelationAccountUser, error) {
	r := &po.RelationAccountUser{
		TenantId: tenantId,
	}
	err := a.postgres.FindOne(r, "account_id=? and deleted=?", accId, 0)
	return r, err
}
