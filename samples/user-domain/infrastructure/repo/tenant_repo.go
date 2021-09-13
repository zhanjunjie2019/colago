package repo

import (
	"github.com/zhanjunjie2019/colago/common/dbcli"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/user-domain/infrastructure/repo/po"
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
