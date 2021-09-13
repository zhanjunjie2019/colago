package user

import (
	"github.com/zhanjunjie2019/colago/common/codec"
	"github.com/zhanjunjie2019/colago/common/domain"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"github.com/zhanjunjie2019/colago/samples/user-domain/domain/account"
	"github.com/zhanjunjie2019/colago/samples/user-domain/domain/auth"
	"time"
)

func init() {
	ioc.AppendPullFactory(func(
		u UserGateway,
		a auth.AuthGateway) {
		userGateway = u
		authGateway = a
	})
}

var (
	userGateway UserGateway
	authGateway auth.AuthGateway
)

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
	dto         *client.DTO
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
		roles, err := authGateway.FindRolesByUserId(u.Ctx(), u.dto, u.Id())
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
		auths, err := authGateway.FindAuthsByUserId(u.Ctx(), u.dto, u.Id())
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
	return userGateway.CreateUser(u.Ctx(), u.dto, u)
}
