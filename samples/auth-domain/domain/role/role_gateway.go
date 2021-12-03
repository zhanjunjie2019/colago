package role

import (
	"context"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
)

type RoleGateway interface {
	FindByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]*Role, error)
}
