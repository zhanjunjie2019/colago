package user

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
)

type UserGateway interface {
	CreateUser(dto *client.DTO, user *User) error
	FindByAccount(dto *client.DTO, acc *account.Account) (*User, error)
}
