package entity

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/common/ioc"
)

func init() {
	_ = ioc.InjectPrototypeBean(new(Account))
}

type AccountType uint8

const (
	Phone AccountType = iota
	Email
	Ordinary
	Channel
)

type Account struct {
	domain.Entity
	accType  AccountType
	accKey   string
	password string
	enable   bool
}

func (a *Account) New() ioc.AbsBean {
	return a
}

func (a *Account) AccType() AccountType {
	return a.accType
}

func (a *Account) SetAccType(accType AccountType) {
	a.accType = accType
}

func (a *Account) AccKey() string {
	return a.accKey
}

func (a *Account) SetAccKey(accKey string) {
	a.accKey = accKey
}

func (a *Account) Password() string {
	return a.password
}

func (a *Account) SetPassword(password string) {
	a.password = password
}

func (a *Account) Enable() bool {
	return a.enable
}

func (a *Account) SetEnable(enable bool) {
	a.enable = enable
}
