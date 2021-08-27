package executor

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/jwt"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/ability"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
)

func init() {
	_ = ioc.InjectSimpleBean(new(UserAppExe))
}

func NewUserAppExe() client.User {
	bean, _ := ioc.GetBean("executor.UserAppExe")
	exe := bean.(*UserAppExe)
	return exe
}

type UserAppExe struct {
	userService *ability.UserService `ij:"ability.UserService"`
}

func (u *UserAppExe) UserService() *ability.UserService {
	return u.userService
}

func (u *UserAppExe) SetUserService(userService *ability.UserService) {
	u.userService = userService
}

func (u *UserAppExe) New() ioc.AbsBean {
	return u
}

func (u *UserAppExe) Init(id string) {
}

func (u *UserAppExe) Terminate() {
}

func (u *UserAppExe) ReceiveDefault(actor.Context) {
}

func (u *UserAppExe) LoginAction(cmd *client.UserLoginCmd, context cluster.GrainContext) (*client.UserLoginResponse, error) {
	dto := cmd.Dto
	key := cmd.AccKey
	password := cmd.Password
	response := new(client.UserLoginResponse)

	token, err := u.userService.LoginAction(dto, key, password)
	if err != nil {
		response.Rsp = &client.Response{
			Success:    false,
			ErrCode:    "ERR001",
			ErrMessage: err.Error(),
		}
	}
	if token != nil {
		response.Rsp = &client.Response{
			Success: true,
		}
		jwtStr, _ := jwt.JwtBuild(*token, "")
		response.Data = &client.UserLoginData{
			Ext:         token.Exp,
			ClientToken: jwtStr,
		}
	}
	return response, nil
}
