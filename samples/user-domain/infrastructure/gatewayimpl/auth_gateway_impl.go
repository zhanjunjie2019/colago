package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	authclient "e.coding.net/double-j/ego/colago/samples/auth-client"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
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

func (a *AuthGatewayImpl) FindRolesByUserId(dto *client.DTO, userId uint64) ([]string, error) {
	return a.Authcli.FindRolesByUserId(&client.RoleQry{
		Dto:    dto,
		UserId: userId,
	})
}

func (a *AuthGatewayImpl) FindAuthsByUserId(dto *client.DTO, userId uint64) ([]string, error) {
	return a.Authcli.FindAuthsByUserId(&client.AuthQry{
		Dto:    dto,
		UserId: userId,
	})
}
