package po

import (
	"github.com/zhanjunjie2019/colago/common/model"
	"strconv"
)

type RelationUserRole struct {
	model.PostgresEntity
	UserId   uint64 `gorm:"comment:用户ID"`
	RoleCode string `gorm:"type:varchar(64);comment:角色编码"`
	TenantId uint64 `gorm:"-"`
}

func (r *RelationUserRole) TableName() string {
	return "t" + strconv.FormatUint(r.TenantId, 10) + "_rela_user_role"
}
