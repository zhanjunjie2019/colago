package entity

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/model/valueobject"
)

func init() {
	_ = ioc.InjectPrototypeBean(new(User))
}

type User struct {
	domain.Entity
	auths []valueobject.Auth
	roles []valueobject.Role
}

func (u *User) New() ioc.AbsBean {
	return u
}

func (u *User) Auths() []valueobject.Auth {
	return u.auths
}

func (u *User) SetAuths(auths []valueobject.Auth) {
	u.auths = auths
}

func (u *User) Roles() []valueobject.Role {
	return u.roles
}

func (u *User) SetRoles(roles []valueobject.Role) {
	u.roles = roles
}
