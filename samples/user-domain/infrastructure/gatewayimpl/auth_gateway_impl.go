package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	authclient "e.coding.net/double-j/ego/colago/samples/auth-client"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"golang.org/x/net/context"
)

func init() {
	err := ioc.InjectSimpleBean(new(AuthGatewayImpl))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type AuthGatewayImpl struct {
	Authcli *authclient.AuthClient `ij:"authclient.AuthClient"`
}

func (a *AuthGatewayImpl) New() ioc.AbsBean {
	return a
}

func (a *AuthGatewayImpl) FindRolesByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]string, error) {
	return a.Authcli.FindRolesByUserId(ctx, &client.RoleQry{
		Dto:    dto,
		UserId: userId,
	})
}

func (a *AuthGatewayImpl) FindAuthsByUserId(ctx context.Context, dto *client.DTO, userId uint64) ([]string, error) {
	return a.Authcli.FindAuthsByUserId(ctx, &client.AuthQry{
		Dto:    dto,
		UserId: userId,
	})
}
