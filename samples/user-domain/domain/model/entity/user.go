package entity

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"time"
)

func init() {
	_ = ioc.InjectPrototypeBean(new(User))
}

type User struct {
	domain.Entity
	accounts    []Account
	firstName   string
	lastName    string
	age         uint8
	birthday    time.Time
	email       string
	phoneNumber string
	status      uint8
	roles       []string
	auths       []string
}

func (u *User) New() ioc.AbsBean {
	return u
}

func (u *User) Accounts() []Account {
	return u.accounts
}

func (u *User) SetAccounts(accounts []Account) {
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
	return u.roles
}

func (u *User) SetRoles(roles []string) {
	u.roles = roles
}

func (u *User) Auths() []string {
	return u.auths
}

func (u *User) SetAuths(auths []string) {
	u.auths = auths
}
