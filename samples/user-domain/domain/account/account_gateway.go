package account

import (
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"golang.org/x/net/context"
)

type AccountGateway interface {
	FindAccountByAccKey(ctx context.Context, dto *client.DTO, accKey string) (*Account, error)
}
