package role

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"golang.org/x/net/context"
)

type RoleGateway interface {
	FindByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]*Role, error)
}
