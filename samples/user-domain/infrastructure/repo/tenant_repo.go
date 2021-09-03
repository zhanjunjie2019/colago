package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(TenantRepo))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type TenantRepo struct {
	Postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (t *TenantRepo) New() ioc.AbsBean {
	return t
}

func (t *TenantRepo) TenantInitAction(tenantId uint64) error {
	return t.Postgres.AutoMigrate(
		&po.Account{
			TenantId: tenantId,
		},
		&po.UserInfo{
			TenantId: tenantId,
		},
		&po.RelationAccountUser{
			TenantId: tenantId,
		},
	)
}
