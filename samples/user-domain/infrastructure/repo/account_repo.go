package repo

import (
	"e.coding.net/double-j/ego/colago/common/dbcli"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func init() {
	ioc.AppendInjection(func(dbcli dbcli.Cli) *AccountRepo {
		return &AccountRepo{
			dbcli: dbcli,
		}
	})
}

type AccountRepo struct {
	dbcli dbcli.Cli
}

func (a *AccountRepo) InsertOne(accPo *po.Account) (*po.Account, error) {
	return accPo, a.dbcli.InsertOne(accPo)
}

func (a *AccountRepo) FindByAccKey(tenantId uint64, acckey string) (*po.Account, error) {
	account := &po.Account{
		TenantId: tenantId,
	}
	err := a.dbcli.FindOne(account, []string{"acc_key=?", "deleted=?"}, []interface{}{acckey, 0})
	return account, err
}
