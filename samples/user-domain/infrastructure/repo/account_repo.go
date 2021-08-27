package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func init() {
	_ = ioc.InjectSimpleBean(new(AccountRepo))
}

type AccountRepo struct {
	postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (a *AccountRepo) Postgres() *postgres.Postgres {
	return a.postgres
}

func (a *AccountRepo) SetPostgres(postgres *postgres.Postgres) {
	a.postgres = postgres
}

func (a *AccountRepo) New() ioc.AbsBean {
	return a
}

func (a *AccountRepo) CreateTable(tenantId uint64) error {
	return a.postgres.AutoMigrate(&po.Account{
		TenantId: tenantId,
	})
}

func (a *AccountRepo) FindByAccKey(tenantId uint64, acckey string) (*po.Account, error) {
	account := &po.Account{
		TenantId: tenantId,
	}
	err := a.postgres.FindOne(account, "acc_key=? and deleted=?", acckey, 0)
	return account, err
}
