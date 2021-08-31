package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo/po"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(UserAuthRepo))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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

func (u *UserAuthRepo) InsertBatch(rePos []*po.RelationUserAuth) ([]*po.RelationUserAuth, error) {
	for _, rePo := range rePos {
		err := u.postgres.InsertOne(rePo)
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
	err := u.postgres.FindList(r, &auths, []string{"user_id=?", "deleted=?"}, []interface{}{userId, 0})
	return auths, err
}
