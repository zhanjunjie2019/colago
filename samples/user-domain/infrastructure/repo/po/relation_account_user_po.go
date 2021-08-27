package po

import (
	"e.coding.net/double-j/ego/colago/common/model"
	"strconv"
)

type RelationAccountUser struct {
	model.PostgresEntity
	AccountId uint64 `gorm:"comment:账号ID"`
	UserId    uint64 `gorm:"comment:用户ID"`
	TenantId  uint64 `gorm:"-"`
}

func (a *RelationAccountUser) TableName() string {
	return "rela_account_user_" + strconv.FormatUint(a.TenantId, 10)
}
