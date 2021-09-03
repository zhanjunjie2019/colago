package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo/po"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(UserRoleRepo))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type UserRoleRepo struct {
	Postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (u *UserRoleRepo) New() ioc.AbsBean {
	return u
}

func (u *UserRoleRepo) InsertBatch(rePos []*po.RelationUserRole) ([]*po.RelationUserRole, error) {
	for _, rePo := range rePos {
		err := u.Postgres.InsertOne(rePo)
		if err != nil {
			return nil, err
		}
	}
	return rePos, nil
}

func (u *UserRoleRepo) ListByUserId(tenantId uint64, userId uint64) ([]*po.RelationUserRole, error) {
	r := &po.RelationUserRole{
		TenantId: tenantId,
	}
	roles := make([]*po.RelationUserRole, 0)
	err := u.Postgres.FindList(r, &roles, []string{"user_id=?", "deleted=?"}, []interface{}{userId, 0})
	return roles, err
}
