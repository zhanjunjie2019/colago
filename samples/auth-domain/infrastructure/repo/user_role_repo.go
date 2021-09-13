package repo

import (
	"e.coding.net/double-j/ego/colago/common/dbcli"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo/po"
)

func init() {
	ioc.AppendInjection(func(dbcli dbcli.Cli) *UserRoleRepo {
		return &UserRoleRepo{
			dbcli: dbcli,
		}
	})
}

type UserRoleRepo struct {
	dbcli dbcli.Cli
}

func (u *UserRoleRepo) InsertBatch(rePos []*po.RelationUserRole) ([]*po.RelationUserRole, error) {
	for _, rePo := range rePos {
		err := u.dbcli.InsertOne(rePo)
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
	err := u.dbcli.FindList(r, &roles, []string{"user_id=?", "deleted=?"}, []interface{}{userId, 0})
	return roles, err
}
