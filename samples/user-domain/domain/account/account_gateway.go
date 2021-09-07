package account

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"golang.org/x/net/context"
)

type AccountGateway interface {
	FindAccountByAccKey(ctx context.Context, dto *client.DTO, accKey string) (*Account, error)
}
