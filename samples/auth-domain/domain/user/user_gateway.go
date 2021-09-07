package user

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"golang.org/x/net/context"
)

type UserGateway interface {
	SaveRoleAuth(ctx context.Context, dto *client.DTO, user *User) error
	FindById(ctx context.Context, dto *client.DTO, userId uint64) (*User, error)
}
