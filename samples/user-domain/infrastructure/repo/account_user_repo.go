package repo

import (
	"github.com/zhanjunjie2019/colago/common/dbcli"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/user-domain/infrastructure/repo/po"
)

func init() {
	ioc.AppendInjection(func(dbcli dbcli.Cli) *AccountUserRepo {
		return &AccountUserRepo{
			dbcli: dbcli,
		}
	})
}

type AccountUserRepo struct {
	dbcli dbcli.Cli
}

func (a *AccountUserRepo) InsertBatch(rePos []*po.RelationAccountUser) ([]*po.RelationAccountUser, error) {
	for _, rePo := range rePos {
		err := a.dbcli.InsertOne(rePo)
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
	err := a.dbcli.FindOne(r, []string{"account_id=?", "deleted=?"}, []interface{}{accId, 0})
	return r, err
}
