package repo

import (
	"github.com/zhanjunjie2019/colago/common/dbcli"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/repo/po"
)

func init() {
	ioc.AppendInjection(func(dbcli dbcli.Cli) *UserAuthRepo {
		return &UserAuthRepo{
			dbcli: dbcli,
		}
	})
}

type UserAuthRepo struct {
	dbcli dbcli.Cli
}

func (u *UserAuthRepo) InsertBatch(rePos []*po.RelationUserAuth) ([]*po.RelationUserAuth, error) {
	for _, rePo := range rePos {
		err := u.dbcli.InsertOne(rePo)
		if err != nil {
			return nil, err
		}
	}
	return rePos, nil
}

func (u *UserAuthRepo) ListByUserId(tenantId uint64, userId uint64) ([]*po.RelationUserAuth, error) {
	r := &po.RelationUserAuth{
		TenantId: tenantId,
	}
	auths := make([]*po.RelationUserAuth, 0)
	err := u.dbcli.FindList(r, &auths, []string{"user_id=?", "deleted=?"}, []interface{}{userId, 0})
	return auths, err
}
