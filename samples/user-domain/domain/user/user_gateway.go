package user

import (
	"context"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"github.com/zhanjunjie2019/colago/samples/user-domain/domain/account"
)

type UserGateway interface {
	CreateUser(ctx context.Context, dto *client.DTO, user *User) error
	FindByAccount(ctx context.Context, dto *client.DTO, acc *account.Account) (*User, error)
}
