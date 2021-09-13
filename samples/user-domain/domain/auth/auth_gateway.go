package auth

import (
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"golang.org/x/net/context"
)

type AuthGateway interface {
	FindRolesByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]string, error)
	FindAuthsByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]string, error)
}
