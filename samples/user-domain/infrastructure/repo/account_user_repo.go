package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(AccountUserRepo))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type AccountUserRepo struct {
	Postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (a *AccountUserRepo) New() ioc.AbsBean {
	return a
}

func (a *AccountUserRepo) InsertBatch(rePos []*po.RelationAccountUser) ([]*po.RelationAccountUser, error) {
	for _, rePo := range rePos {
		err := a.Postgres.InsertOne(rePo)
		if err != nil {
			return nil, err
		}
	}
	return rePos, nil
}

func (a *AccountUserRepo) FindByAccountId(tenantId uint64, accId uint64) (*po.RelationAccountUser, error) {
	r := &po.RelationAccountUser{
		TenantId: tenantId,
	}
	err := a.Postgres.FindOne(r, []string{"account_id=?", "deleted=?"}, []interface{}{accId, 0})
	return r, err
}
