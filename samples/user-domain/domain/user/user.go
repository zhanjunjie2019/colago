package user

import (
	"e.coding.net/double-j/ego/colago/common/codec"
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/auth"
	"fmt"
	"time"
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
	accounts    []*account.Account
	firstName   string
	lastName    string
	age         uint8
	birthday    time.Time
	email       string
	phoneNumber string
	status      uint8
	roles       []string
	auths       []string
	userGateway UserGateway      `ij:"gatewayimpl.UserGatewayImpl"`
	authGateway auth.AuthGateway `ij:"gatewayimpl.AuthGatewayImpl"`
	dto         *client.DTO
}

func (u *User) New() ioc.AbsBean {
	return u
}

func (u *User) Accounts() []*account.Account {
	return u.accounts
}

func (u *User) SetAccounts(accounts []*account.Account) {
	u.accounts = accounts
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) Age() uint8 {
	return u.age
}

func (u *User) SetAge(age uint8) {
	u.age = age
}

func (u *User) Birthday() time.Time {
	return u.birthday
}

func (u *User) SetBirthday(birthday time.Time) {
	u.birthday = birthday
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) PhoneNumber() string {
	return u.phoneNumber
}

func (u *User) SetPhoneNumber(phoneNumber string) {
	u.phoneNumber = phoneNumber
}

func (u *User) Status() uint8 {
	return u.status
}

func (u *User) SetStatus(status uint8) {
	u.status = status
}

func (u *User) Roles() []string {
	if u.roles == nil {
		roles, err := u.authGateway.FindRolesByUserId(u.dto, u.Id())
		if err != nil {
			// TODO 还不知道怎么办
			return nil
		}
		u.roles = roles
	}
	return u.roles
}

func (u *User) SetRoles(roles []string) {
	u.roles = roles
}

func (u *User) Auths() []string {
	if u.auths == nil {
		auths, err := u.authGateway.FindAuthsByUserId(u.dto, u.Id())
		if err != nil {
			// TODO 还不知道怎么办
			return nil
		}
		u.auths = auths
	}
	return u.auths
}

func (u *User) SetAuths(auths []string) {
	u.auths = auths
}

func (u *User) UserGateway() UserGateway {
	return u.userGateway
}

func (u *User) SetUserGateway(userGateway UserGateway) {
	u.userGateway = userGateway
}

func (u *User) AuthGateway() auth.AuthGateway {
	return u.authGateway
}

func (u *User) SetAuthGateway(authActuator auth.AuthGateway) {
	u.authGateway = authActuator
}

func (u *User) Dto() *client.DTO {
	return u.dto
}

func (u *User) SetDto(dto *client.DTO) {
	u.dto = dto
}

func (u *User) Create() error {
	for _, a := range u.accounts {
		a.SetPassword(codec.ToSHA1(a.Password()))
	}
	return u.userGateway.CreateUser(u.dto, u)
}
