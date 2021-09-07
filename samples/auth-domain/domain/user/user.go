package user

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/auth"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/role"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
)

func init() {
	err := ioc.InjectPrototypeBean(new(User))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type User struct {
	domain.Entity
	auths       []*auth.Auth
	roles       []*role.Role
	UserGateway UserGateway      `ij:"gatewayimpl.UserGatewayImpl"`
	RoleGateway role.RoleGateway `ij:"gatewayimpl.RoleGatewayImpl"`
	AuthGateway auth.AuthGateway `ij:"gatewayimpl.AuthGatewayImpl"`
	dto         *client.DTO
}

func (u *User) New() ioc.AbsBean {
	return u
}

func (u *User) Auths() []*auth.Auth {
	if u.auths == nil {
		auths, err := u.AuthGateway.FindByUserId(u.Ctx(), u.dto, u.Id())
		if err != nil {
			// TODO 还不知道怎么办
			return nil
		}
		u.auths = auths
	}
	return u.auths
}

func (u *User) SetAuths(auths []*auth.Auth) {
	u.auths = auths
}

func (u *User) Roles() []*role.Role {
	if u.roles == nil {
		roles, err := u.RoleGateway.FindByUserId(u.Ctx(), u.dto, u.Id())
		if err != nil {
			// TODO 还不知道怎么办
			return nil
		}
		u.roles = roles
	}
	return u.roles
}

func (u *User) SetRoles(roles []*role.Role) {
	u.roles = roles
}

func (u *User) Dto() *client.DTO {
	return u.dto
}

func (u *User) SetDto(dto *client.DTO) {
	u.dto = dto
}

func (u *User) SaveRoleAuth() error {
	return u.UserGateway.SaveRoleAuth(u.Ctx(), u.dto, u)
}
