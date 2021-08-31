package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(AccountRepo))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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

func (a *AccountRepo) InsertOne(accPo *po.Account) (*po.Account, error) {
	return accPo, a.postgres.InsertOne(accPo)
}

func (a *AccountRepo) FindByAccKey(tenantId uint64, acckey string) (*po.Account, error) {
	account := &po.Account{
		TenantId: tenantId,
	}
	err := a.postgres.FindOne(account, []string{"acc_key=?", "deleted=?"}, []interface{}{acckey, 0})
	return account, err
}
