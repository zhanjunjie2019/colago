package repo

import (
	"e.coding.net/double-j/ego/colago/common/dbcli"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func init() {
	ioc.AppendInjection(func(dbcli dbcli.Cli) *UserRepo {
		return &UserRepo{
			dbcli: dbcli,
		}
	})
}

type UserRepo struct {
	dbcli dbcli.Cli
}

func (u *UserRepo) InsertOne(userPo *po.UserInfo) (*po.UserInfo, error) {
	return userPo, u.dbcli.InsertOne(userPo)
}

func (u *UserRepo) FindById(tenantId uint64, userId uint64) (*po.UserInfo, error) {
	user := &po.UserInfo{
		TenantId: tenantId,
	}
	err := u.dbcli.FindOne(user, []string{"id=?", "deleted=?"}, []interface{}{userId, 0})
	return user, err
}
