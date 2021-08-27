package model

import "encoding/json"

type TokenData struct {
	// 租户ID
	TenantId uint64
	// 用户ID
	UserId uint64
	// 用户名
	UserName string
	// 角色集
	HasRoles []string
	// 权限
	HasAuths []string
	// 时效
	Exp uint64
}

func (tokenData TokenData) MarshalBinary() (data []byte, err error) {
	return json.Marshal(tokenData)
}
