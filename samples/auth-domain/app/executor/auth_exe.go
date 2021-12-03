package executor

import (
	"context"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/domain/user"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/convertor"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/repo"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
)

var authAppExe *AuthAppExe

func NewAuthAppExe() client.Auth {
	if authAppExe == nil {
		authAppExe = new(AuthAppExe)
		err := ioc.GetContainer().Invoke(func(
			userGateway user.UserGateway,
			tenantRepo *repo.TenantRepo) {
			authAppExe.userGateway = userGateway
			authAppExe.tenantRepo = tenantRepo
		})
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
	}
	return authAppExe
}

type AuthAppExe struct {
	userGateway user.UserGateway
	tenantRepo  *repo.TenantRepo
}

func (a *AuthAppExe) Init(string) {
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

func (a *AuthAppExe) CreateAuthAction(cmd *client.CreateAuthCmd, ctx cluster.GrainContext) (*client.AuthResponse, error) {
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
	u, err := a.userGateway.FindById(context.Background(), dto, userid)
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

func (a *AuthAppExe) FindRolesByUserId(qry *client.RoleQry, ctx cluster.GrainContext) (*client.RoleQryResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dto := qry.Dto
	userid := qry.UserId
	response := new(client.RoleQryResponse)
	u, err := a.userGateway.FindById(context.Background(), dto, userid)
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

func (a *AuthAppExe) FindAuthsByUserId(qry *client.AuthQry, ctx cluster.GrainContext) (*client.AuthQryResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dto := qry.Dto
	userid := qry.UserId
	response := new(client.AuthQryResponse)
	u, err := a.userGateway.FindById(context.Background(), dto, userid)
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
