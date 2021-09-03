package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(UserRepo))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type UserRepo struct {
	Postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (u *UserRepo) New() ioc.AbsBean {
	return u
}

func (u *UserRepo) InsertOne(userPo *po.UserInfo) (*po.UserInfo, error) {
	return userPo, u.Postgres.InsertOne(userPo)
}

func (u *UserRepo) FindById(tenantId uint64, userId uint64) (*po.UserInfo, error) {
	user := &po.UserInfo{
		TenantId: tenantId,
	}
	err := u.Postgres.FindOne(user, []string{"id=?", "deleted=?"}, []interface{}{userId, 0})
	return user, err
}
