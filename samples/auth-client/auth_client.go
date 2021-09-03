package authclient

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/common/sentinel"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"strconv"
	"time"
)

func init() {
	err := ioc.InjectSimpleBean(new(AuthClient))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type AuthClient struct {
	Sent *sentinel.Sentinel `ij:"sentinel.Sentinel"`
}

func (a *AuthClient) New() ioc.AbsBean {
	a.Sent.AppendCircuitbreakerRules(
		&circuitbreaker.Rule{
			Resource:         "Auth.InitAuthTenant",
			Strategy:         circuitbreaker.ErrorCount, // 异常记数方案
			RetryTimeoutMs:   3000,                      // 熔断后3秒重试
			MinRequestAmount: 10,                        // 单位时间内10个请求以上才进入异常记数计算
			StatIntervalMs:   5000,                      // 单位时间为5秒
			Threshold:        10,                        // 单位时间内容错数量
		},
		&circuitbreaker.Rule{
			Resource:         "Auth.CreateRoleAuthCodes",
			Strategy:         circuitbreaker.ErrorCount, // 异常记数方案
			RetryTimeoutMs:   3000,                      // 熔断后3秒重试
			MinRequestAmount: 10,                        // 单位时间内10个请求以上才进入异常记数计算
			StatIntervalMs:   5000,                      // 单位时间为5秒
			Threshold:        10,                        // 单位时间内容错数量
		},
		&circuitbreaker.Rule{
			Resource:         "Auth.FindRolesByUserId",
			Strategy:         circuitbreaker.ErrorCount, // 异常记数方案
			RetryTimeoutMs:   3000,                      // 熔断后3秒重试
			MinRequestAmount: 10,                        // 单位时间内10个请求以上才进入异常记数计算
			StatIntervalMs:   5000,                      // 单位时间为5秒
			Threshold:        10,                        // 单位时间内容错数量
		},
		&circuitbreaker.Rule{
			Resource:         "Auth.FindAuthsByUserId",
			Strategy:         circuitbreaker.ErrorCount, // 异常记数方案
			RetryTimeoutMs:   3000,                      // 熔断后3秒重试
			MinRequestAmount: 10,                        // 单位时间内10个请求以上才进入异常记数计算
			StatIntervalMs:   5000,                      // 单位时间为5秒
			Threshold:        10,                        // 单位时间内容错数量
		},
	)
	a.Sent.LoadRules()
	return a
}

func (a *AuthClient) InitAuthTenant(dto *client.AuthTenantInitCmd) error {
	_, err := protoactor.ClientChain(
		protoactor.ClientActionArgs{
			Resource: "Auth.InitAuthTenant",
			TryFn: func() (interface{}, error) {
				callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
				grainClient := client.GetAuthGrainClient(protoactor.Cluster, strconv.FormatUint(dto.TenantId, 10))
				authResponse, err := grainClient.TenantInitAction(dto, callOpts)
				if err != nil {
					panic(err)
				}
				if !authResponse.Rsp.Success {
					return nil, fmt.Errorf("[AUTH:" + authResponse.Rsp.ErrCode + ":" + authResponse.Rsp.ErrMessage + "]")
				} else {
					return nil, nil
				}
			},
			CatchFn: func(er interface{}) (interface{}, error) {
				er2, ok := er.(error)
				if ok {
					return nil, er2
				}
				return nil, fmt.Errorf("%v", er)
			},
		},
	)
	return err
}

func (a *AuthClient) CreateRoleAuthCodes(dto *client.CreateAuthCmd) error {
	_, err := protoactor.ClientChain(
		protoactor.ClientActionArgs{
			Resource: "Auth.CreateRoleAuthCodes",
			TryFn: func() (interface{}, error) {
				callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
				grainClient := client.GetAuthGrainClient(protoactor.Cluster, strconv.FormatUint(dto.UserId, 10))
				authResponse, err := grainClient.CreateAuthAction(dto, callOpts)
				if err != nil {
					panic(err)
				}
				if !authResponse.Rsp.Success {
					return nil, fmt.Errorf("[AUTH:" + authResponse.Rsp.ErrCode + ":" + authResponse.Rsp.ErrMessage + "]")
				} else {
					return nil, nil
				}
			},
			CatchFn: func(er interface{}) (interface{}, error) {
				er2, ok := er.(error)
				if ok {
					return nil, er2
				}
				return nil, fmt.Errorf("%v", er)
			},
		},
	)
	return err
}

func (a *AuthClient) FindRolesByUserId(dto *client.RoleQry) ([]string, error) {
	rs, err := protoactor.ClientChain(
		protoactor.ClientActionArgs{
			Resource: "Auth.FindRolesByUserId",
			TryFn: func() (interface{}, error) {
				callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
				grainClient := client.GetAuthGrainClient(protoactor.Cluster, strconv.FormatUint(dto.UserId, 10))
				roleQryResponse, err := grainClient.FindRolesByUserId(dto, callOpts)
				if err != nil {
					panic(err)
				}
				if !roleQryResponse.Rsp.Success {
					return nil, fmt.Errorf("[AUTH:" + roleQryResponse.Rsp.ErrCode + ":" + roleQryResponse.Rsp.ErrMessage + "]")
				} else {
					return roleQryResponse.Roles, nil
				}
			},
			CatchFn: func(er interface{}) (interface{}, error) {
				er2, ok := er.(error)
				if ok {
					return nil, er2
				}
				return nil, fmt.Errorf("%v", er)
			},
		},
	)
	if rs != nil {
		return rs.([]string), nil
	}
	return nil, err
}

func (a *AuthClient) FindAuthsByUserId(dto *client.AuthQry) ([]string, error) {
	rs, err := protoactor.ClientChain(
		protoactor.ClientActionArgs{
			Resource: "Auth.FindAuthsByUserId",
			TryFn: func() (interface{}, error) {
				callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
				grainClient := client.GetAuthGrainClient(protoactor.Cluster, strconv.FormatUint(dto.UserId, 10))
				authQryResponse, err := grainClient.FindAuthsByUserId(dto, callOpts)
				if err != nil {
					panic(err)
				}
				if !authQryResponse.Rsp.Success {
					return nil, fmt.Errorf("[AUTH:" + authQryResponse.Rsp.ErrCode + ":" + authQryResponse.Rsp.ErrMessage + "]")
				} else {
					return authQryResponse.Auths, nil
				}
			},
			CatchFn: func(er interface{}) (interface{}, error) {
				er2, ok := er.(error)
				if ok {
					return nil, er2
				}
				return nil, fmt.Errorf("%v", er)
			},
		},
	)
	if rs != nil {
		return rs.([]string), nil
	}
	return nil, err
}
