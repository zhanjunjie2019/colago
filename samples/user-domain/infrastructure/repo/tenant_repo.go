package repo

import (
	"e.coding.net/double-j/ego/colago/common/dbcli"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func init() {
	ioc.AppendInjection(func(dbcli dbcli.Cli) *TenantRepo {
		return &TenantRepo{
			dbcli: dbcli,
		}
	})
}

type TenantRepo struct {
	dbcli dbcli.Cli
}

func (t *TenantRepo) TenantInitAction(tenantId uint64) error {
	return t.dbcli.AutoMigrate(
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
