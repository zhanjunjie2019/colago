package executor

import (
	"context"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/common/jwt"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"github.com/zhanjunjie2019/colago/samples/user-domain/domain/user"
	"github.com/zhanjunjie2019/colago/samples/user-domain/infrastructure/convertor"
	"github.com/zhanjunjie2019/colago/samples/user-domain/infrastructure/repo"
)

var userAppExe *UserAppExe

func NewUserAppExe() client.User {
	if userAppExe == nil {
		userAppExe = new(UserAppExe)
		err := ioc.GetContainer().Invoke(func(
			userService *user.UserService,
			tenantRepo *repo.TenantRepo) {
			userAppExe.userService = userService
			userAppExe.tenantRepo = tenantRepo
		})
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
	}
	return userAppExe
}

type UserAppExe struct {
	userService *user.UserService
	tenantRepo  *repo.TenantRepo
}

func (u *UserAppExe) Init(string) {
}

func (u *UserAppExe) Terminate() {
}

func (u *UserAppExe) ReceiveDefault(actor.Context) {
}

func (u *UserAppExe) TenantInitAction(cmd *client.UserTenantInitCmd, context cluster.GrainContext) (*client.UserResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	response := new(client.UserResponse)
	err := u.tenantRepo.TenantInitAction(cmd.TenantId)
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

func (u *UserAppExe) CreateUserAction(cmd *client.CreateUserCmd, ctx cluster.GrainContext) (*client.UserResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	response := new(client.UserResponse)
	entity, err := convertor.UserCreateDtoToUserEntity(context.Background(), cmd)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR002",
			ErrMessage: err.Error(),
		}
		return response, nil
	}
	if entity != nil {
		err = entity.Create()
		if err != nil {
			response.Rsp = &client.Response{
				Success:    false,
				ErrCode:    "ERR003",
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

func (u *UserAppExe) LoginAction(cmd *client.UserLoginCmd, ctx cluster.GrainContext) (*client.UserLoginResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dto := cmd.Dto
	key := cmd.AccKey
	password := cmd.Password
	response := new(client.UserLoginResponse)
	token, err := u.userService.LoginAction(context.Background(), dto, key, password)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR004",
			ErrMessage: err.Error(),
		}
		return response, nil
	}
	if token != nil {
		response.Rsp = &client.Response{
			Success: true,
		}
		jwtStr, _ := jwt.JwtBuild(*token, "abcdefg")
		response.Data = &client.UserLoginData{
			Ext:         token.Exp,
			ClientToken: jwtStr,
		}
	}
	return response, nil
}
