package user_client

import (
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"strconv"
	"time"
)

func InitUserTenant(dto *client.UserTenantInitCmd) error {
	callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
	grainClient := client.GetUserGrainClient(protoactor.Cluster, strconv.FormatUint(dto.TenantId, 10))
	action, err := grainClient.TenantInitAction(dto, callOpts)
	if err != nil {
		return err
	}
	if !action.Rsp.Success {
		return fmt.Errorf("[USER:" + action.Rsp.ErrCode + ":" + action.Rsp.ErrMessage + "]")
	} else {
		return nil
	}
}

func CreateUserAction(dto *client.CreateUserCmd) error {
	callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
	grainClient := client.GetUserGrainClient(protoactor.Cluster, dto.AccKey)
	action, err := grainClient.CreateUserAction(dto, callOpts)
	if err != nil {
		return err
	}
	if !action.Rsp.Success {
		return fmt.Errorf("[USER:" + action.Rsp.ErrCode + ":" + action.Rsp.ErrMessage + "]")
	} else {
		return nil
	}
}

func LoginAction(dto *client.UserLoginCmd) (*client.UserLoginData, error) {
	callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
	grainClient := client.GetUserGrainClient(protoactor.Cluster, dto.AccKey)
	action, err := grainClient.LoginAction(dto, callOpts)
	if err != nil {
		return nil, err
	}
	if !action.Rsp.Success {
		return nil, fmt.Errorf("[USER:" + action.Rsp.ErrCode + ":" + action.Rsp.ErrMessage + "]")
	} else {
		return action.Data, nil
	}
}
