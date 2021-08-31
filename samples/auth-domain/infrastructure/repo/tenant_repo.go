package repo

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/postgres"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo/po"
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
	postgres *postgres.Postgres `ij:"postgres.Postgres"`
}

func (t *TenantRepo) Postgres() *postgres.Postgres {
	return t.postgres
}

func (t *TenantRepo) SetPostgres(postgres *postgres.Postgres) {
	t.postgres = postgres
}

func (t *TenantRepo) New() ioc.AbsBean {
	return t
}

func (t *TenantRepo) TenantInitAction(tenantId uint64) error {
	return t.postgres.AutoMigrate(
		&po.RelationUserAuth{
			TenantId: tenantId,
		},
		&po.RelationUserRole{
			TenantId: tenantId,
		},
	)
}
