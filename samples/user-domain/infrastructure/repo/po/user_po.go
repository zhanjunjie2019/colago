package po

import (
	"e.coding.net/double-j/ego/colago/common/model"
	"strconv"
	"time"
)

type UserInfo struct {
	model.PostgresEntity
	FirstName   string    `gorm:"type:varchar(256);comment:姓氏"`
	LastName    string    `gorm:"type:varchar(256);comment:名称"`
	Age         uint8     `gorm:"type:int4;comment:年龄"`
	Birthday    time.Time `gorm:"type:date;comment:出生日期"`
	Email       string    `gorm:"type:varchar(256);comment:邮箱"`
	PhoneNumber string    `gorm:"type:varchar(256);comment:收集号码"`
	Status      uint8     `gorm:"type:int2;default:1;comment:状态（0：禁用，1：启用，2：封锁，3：异常审核）"`
	TenantId    uint64    `gorm:"-"`
}

func (u *UserInfo) TableName() string {
	return "t" + strconv.FormatUint(u.TenantId, 10) + "_userinfos"
}
