package gatewayimpl

import (
	"github.com/zhanjunjie2019/colago/common/ioc"
	authclient "github.com/zhanjunjie2019/colago/samples/auth-client"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"github.com/zhanjunjie2019/colago/samples/user-domain/domain/auth"
	"golang.org/x/net/context"
)

func init() {
	ioc.AppendInjection(func(authcli *authclient.AuthClient) auth.AuthGateway {
		return &AuthGatewayImpl{
			authcli: authcli,
		}
	})
}

type AuthGatewayImpl struct {
	authcli *authclient.AuthClient
}

func (a *AuthGatewayImpl) FindRolesByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]string, error) {
	return a.authcli.FindRolesByUserId(ctx, &client.RoleQry{
		Dto:    dto,
		UserId: userId,
	})
}

func (a *AuthGatewayImpl) FindAuthsByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]string, error) {
	return a.authcli.FindAuthsByUserId(ctx, &client.AuthQry{
		Dto:    dto,
		UserId: userId,
	})
}
