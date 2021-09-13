package role

import (
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"golang.org/x/net/context"
)

type RoleGateway interface {
	FindByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]*Role, error)
}
