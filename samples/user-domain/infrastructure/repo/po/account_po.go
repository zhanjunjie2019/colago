package po

import (
	"e.coding.net/double-j/ego/colago/common/model"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"strconv"
)

type Account struct {
	model.PostgresEntity
	AccType  account.AccountType `gorm:"type:int4;comment:账号类型"`
	AccKey   string              `gorm:"type:varchar(256);comment:账号"`
	Password string              `gorm:"type:varchar(256);comment:密码密文"`
	Enable   uint8               `gorm:"type:int2;default:1;comment:启用状态"`
	TenantId uint64              `gorm:"-"`
}

func (a *Account) TableName() string {
	return "t" + strconv.FormatUint(a.TenantId, 10) + "_accounts"
}
