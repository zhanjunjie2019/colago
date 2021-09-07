package user

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"golang.org/x/net/context"
)

type UserGateway interface {
	CreateUser(ctx context.Context, dto *client.DTO, user *User) error
	FindByAccount(ctx context.Context, dto *client.DTO, acc *account.Account) (*User, error)
}
