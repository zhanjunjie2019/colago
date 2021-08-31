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
	return "t" + strconv.FormatUint(a.TenantId, 10) + "_rela_account_user"
}
