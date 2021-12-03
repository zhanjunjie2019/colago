package auth

import (
	"context"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
)

type AuthGateway interface {
	FindByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]*Auth, error)
}
