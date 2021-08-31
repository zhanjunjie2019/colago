package executor

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/user"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
)

func init() {
	err := ioc.InjectSimpleBean(new(AuthAppExe))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func NewAuthAppExe() client.Auth {
	bean, err := ioc.GetBean("executor.AuthAppExe")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	exe := bean.(*AuthAppExe)
	return exe
}

type AuthAppExe struct {
	userGateway user.UserGateway `ij:"gatewayimpl.UserGatewayImpl"`
	tenantRepo  *repo.TenantRepo `ij:"repo.TenantRepo"`
}

func (a *AuthAppExe) TenantRepo() *repo.TenantRepo {
	return a.tenantRepo
}

func (a *AuthAppExe) SetTenantRepo(tenantRepo *repo.TenantRepo) {
	a.tenantRepo = tenantRepo
}

func (a *AuthAppExe) UserGateway() user.UserGateway {
	return a.userGateway
}

func (a *AuthAppExe) SetUserGateway(userGateway user.UserGateway) {
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

func (a *AuthAppExe) TenantInitAction(cmd *client.AuthTenantInitCmd, context cluster.GrainContext) (*client.AuthResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	response := new(client.AuthResponse)
	err := a.tenantRepo.TenantInitAction(cmd.TenantId)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR001",
			ErrMessage: err.Error(),
		}
		return response, nil
	}
	response.Rsp = &client.Response{
		Success: true,
	}

	return response, nil
}

func (a *AuthAppExe) CreateAuthAction(cmd *client.CreateAuthCmd, context cluster.GrainContext) (*client.AuthResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dto := cmd.Dto
	userid := cmd.UserId
	roles := cmd.Roles
	auths := cmd.Auths
	response := new(client.AuthResponse)
	u, err := a.userGateway.FindById(dto, userid)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR001",
			ErrMessage: err.Error(),
		}
		return response, nil
	}
	if u != nil {
		roleEntitys, err := convertor.CodesToBatchRoleEntitys(roles)
		if err != nil {
			response.Rsp = &client.Response{
				Success:    false,
				ErrCode:    "ERR001",
				ErrMessage: err.Error(),
			}
			return response, nil
		}
		u.SetRoles(roleEntitys)
		authEntitys, err := convertor.CodesToBatchAuthEntitys(auths)
		if err != nil {
			response.Rsp = &client.Response{
				Success:    false,
				ErrCode:    "ERR001",
				ErrMessage: err.Error(),
			}
			return response, nil
		}
		u.SetAuths(authEntitys)
		err = u.SaveRoleAuth()
		if err != nil {
			response.Rsp = &client.Response{
				Success:    false,
				ErrCode:    "ERR001",
				ErrMessage: err.Error(),
			}
			return response, nil
		}
		response.Rsp = &client.Response{
			Success: true,
		}
	}
	return response, nil
}

func (a *AuthAppExe) FindRolesByUserId(qry *client.RoleQry, context cluster.GrainContext) (*client.RoleQryResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dto := qry.Dto
	userid := qry.UserId
	response := new(client.RoleQryResponse)
	u, err := a.userGateway.FindById(dto, userid)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR001",
			ErrMessage: err.Error(),
		}
		return response, nil
	}
	if u != nil {
		response.Rsp = &client.Response{
			Success: true,
		}
		roleCodes := make([]string, 0)
		for _, v := range u.Roles() {
			roleCodes = append(roleCodes, v.Code())
		}
		response.Roles = roleCodes
	}
	return response, nil
}

func (a *AuthAppExe) FindAuthsByUserId(qry *client.AuthQry, context cluster.GrainContext) (*client.AuthQryResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dto := qry.Dto
	userid := qry.UserId
	response := new(client.AuthQryResponse)
	u, err := a.userGateway.FindById(dto, userid)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR001",
			ErrMessage: err.Error(),
		}
		return response, nil
	}
	if u != nil {
		response.Rsp = &client.Response{
			Success: true,
		}
		authCodes := make([]string, 0)
		for _, v := range u.Auths() {
			authCodes = append(authCodes, v.Code())
		}
		response.Auths = authCodes
	}
	return response, nil
}
