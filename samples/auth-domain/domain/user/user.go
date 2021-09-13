package user

import (
	"github.com/zhanjunjie2019/colago/common/domain"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/domain/auth"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/domain/role"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
)

func init() {
	ioc.AppendPullFactory(func(
		u UserGateway,
		r role.RoleGateway,
		a auth.AuthGateway) {
		userGateway = u
		roleGateway = r
		authGateway = a
	})
}

var (
	userGateway UserGateway
	roleGateway role.RoleGateway
	authGateway auth.AuthGateway
)

type User struct {
	domain.Entity
	auths []*auth.Auth
	roles []*role.Role
	dto   *client.DTO
}

func (u *User) Auths() []*auth.Auth {
	if u.auths == nil {
		auths, err := authGateway.FindByUserId(u.Ctx(), u.dto, u.Id())
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
		roles, err := roleGateway.FindByUserId(u.Ctx(), u.dto, u.Id())
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
	return userGateway.SaveRoleAuth(u.Ctx(), u.dto, u)
}
