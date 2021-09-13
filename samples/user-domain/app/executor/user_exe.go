package executor

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/jwt"
	"e.coding.net/double-j/ego/colago/common/skywalking"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/user"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
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
	span, err := skywalking.NewRootSpan("User", func(key string) (string, error) {
		return cmd.Dto.TraceId, nil
	})
	defer func() {
		span.End(err)
	}()
	response := new(client.UserResponse)
	err = u.tenantRepo.TenantInitAction(cmd.TenantId)
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

func (u *UserAppExe) CreateUserAction(cmd *client.CreateUserCmd, context cluster.GrainContext) (*client.UserResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	span, err := skywalking.NewRootSpan("User", func(key string) (string, error) {
		return cmd.Dto.TraceId, nil
	})
	defer func() {
		span.End(err)
	}()
	response := new(client.UserResponse)
	entity, err := convertor.UserCreateDtoToUserEntity(span.Ctx(), cmd)
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

func (u *UserAppExe) LoginAction(cmd *client.UserLoginCmd, context cluster.GrainContext) (*client.UserLoginResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	span, err := skywalking.NewRootSpan("User", func(key string) (string, error) {
		return cmd.Dto.TraceId, nil
	})
	defer func() {
		span.End(err)
	}()
	dto := cmd.Dto
	key := cmd.AccKey
	password := cmd.Password
	response := new(client.UserLoginResponse)
	token, err := u.userService.LoginAction(span.Ctx(), dto, key, password)
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
