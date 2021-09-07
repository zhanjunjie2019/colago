package auth

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"golang.org/x/net/context"
)

type AuthGateway interface {
	FindByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]*Auth, error)
}
