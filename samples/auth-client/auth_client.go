package auth_client

import (
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"strconv"
	"time"
)

func InitAuthTenant(dto *client.AuthTenantInitCmd) error {
	callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
	grainClient := client.GetAuthGrainClient(protoactor.Cluster, strconv.FormatUint(dto.TenantId, 10))
	authResponse, err := grainClient.TenantInitAction(dto, callOpts)
	if err != nil {
		return err
	}
	if !authResponse.Rsp.Success {
		return fmt.Errorf("[AUTH:" + authResponse.Rsp.ErrCode + ":" + authResponse.Rsp.ErrMessage + "]")
	} else {
		return nil
	}
}

func CreateRoleAuthCodes(dto *client.CreateAuthCmd) error {
	callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
	grainClient := client.GetAuthGrainClient(protoactor.Cluster, strconv.FormatUint(dto.UserId, 10))
	authResponse, err := grainClient.CreateAuthAction(dto, callOpts)
	if err != nil {
		return err
	}
	if !authResponse.Rsp.Success {
		return fmt.Errorf("[AUTH:" + authResponse.Rsp.ErrCode + ":" + authResponse.Rsp.ErrMessage + "]")
	} else {
		return nil
	}
}

func FindRolesByUserId(dto *client.RoleQry) ([]string, error) {
	callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
	grainClient := client.GetAuthGrainClient(protoactor.Cluster, strconv.FormatUint(dto.UserId, 10))
	roleQryResponse, err := grainClient.FindRolesByUserId(dto, callOpts)
	if err != nil {
		return nil, err
	}
	if !roleQryResponse.Rsp.Success {
		return nil, fmt.Errorf("[AUTH:" + roleQryResponse.Rsp.ErrCode + ":" + roleQryResponse.Rsp.ErrMessage + "]")
	} else {
		return roleQryResponse.Roles, nil
	}
}

func FindAuthsByUserId(dto *client.AuthQry) ([]string, error) {
	callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
	grainClient := client.GetAuthGrainClient(protoactor.Cluster, strconv.FormatUint(dto.UserId, 10))
	authQryResponse, err := grainClient.FindAuthsByUserId(dto, callOpts)
	if err != nil {
		return nil, err
	}
	if !authQryResponse.Rsp.Success {
		return nil, fmt.Errorf("[AUTH:" + authQryResponse.Rsp.ErrCode + ":" + authQryResponse.Rsp.ErrMessage + "]")
	} else {
		return authQryResponse.Auths, nil
	}
}
