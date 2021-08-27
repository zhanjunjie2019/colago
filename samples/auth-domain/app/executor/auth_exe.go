package executor

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/gateway"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
)

func NewAuthAppExe() client.Auth {
	bean, _ := ioc.GetBean("executor.AuthAppExe")
	exe := bean.(*AuthAppExe)
	return exe
}

type AuthAppExe struct {
	userGateway gateway.UserGateway `ij:"gatewayimpl.UserGatewayImpl"`
}

func (a *AuthAppExe) UserGateway() gateway.UserGateway {
	return a.userGateway
}

func (a *AuthAppExe) SetUserGateway(userGateway gateway.UserGateway) {
	a.userGateway = userGateway
}

func (a *AuthAppExe) New() ioc.AbsBean {
	return a
}

func (a *AuthAppExe) Init(id string) {
}

func (a *AuthAppExe) Terminate() {
}

func (a *AuthAppExe) ReceiveDefault(ctx actor.Context) {
}

func (a *AuthAppExe) FindRolesByUserId(qry *client.RoleQry, context cluster.GrainContext) (*client.RoleQryResponse, error) {
	dto := qry.Dto
	userid := qry.UserId
	response := new(client.RoleQryResponse)
	user, err := a.userGateway.FindById(dto, userid)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR001",
			ErrMessage: err.Error(),
		}
	}
	if user != nil {
		response.Rsp = &client.Response{
			Success: true,
		}
		roleCodes := make([]string, 0)
		for _, v := range user.Roles() {
			roleCodes = append(roleCodes, v.Code())
		}
		response.Roles = roleCodes
	}
	return response, nil
}

func (a *AuthAppExe) FindAuthsByUserId(qry *client.AuthQry, context cluster.GrainContext) (*client.AuthQryResponse, error) {
	dto := qry.Dto
	userid := qry.UserId
	response := new(client.AuthQryResponse)
	user, err := a.userGateway.FindById(dto, userid)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR001",
			ErrMessage: err.Error(),
		}
	}
	if user != nil {
		response.Rsp = &client.Response{
			Success: true,
		}
		authCodes := make([]string, 0)
		for _, v := range user.Auths() {
			authCodes = append(authCodes, v.Code())
		}
		response.Auths = authCodes
	}
	return response, nil
}
