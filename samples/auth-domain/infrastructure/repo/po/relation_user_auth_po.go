package po

import (
	"e.coding.net/double-j/ego/colago/common/model"
	"strconv"
)

type RelationUserAuth struct {
	model.PostgresEntity
	UserId   uint64 `gorm:"comment:用户ID"`
	AuthCode string `gorm:"type:varchar(64);comment:权限编码"`
	TenantId uint64 `gorm:"-"`
}

func (r *RelationUserAuth) TableName() string {
	return "t" + strconv.FormatUint(r.TenantId, 10) + "_rela_user_auth"
}
