package auth

import (
	"context"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
)

type AuthGateway interface {
	FindRolesByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]string, error)
	FindAuthsByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]string, error)
}
