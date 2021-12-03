package account

import (
	"context"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
)

type AccountGateway interface {
	FindAccountByAccKey(ctx context.Context, dto *client.DTO, accKey string) (*Account, error)
}
