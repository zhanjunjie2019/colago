package userclient

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/common/protoactor"
	"github.com/zhanjunjie2019/colago/common/sentinel"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func init() {
	ioc.AppendInjection(func(sent *sentinel.Sentinel) *UserClient {
		sent.AppendCircuitbreakerRules(
			&circuitbreaker.Rule{
				Resource:         "User",
				Strategy:         circuitbreaker.ErrorCount, // 异常记数方案
				RetryTimeoutMs:   3000,                      // 熔断后3秒重试
				MinRequestAmount: 10,                        // 单位时间内10个请求以上才进入异常记数计算
				StatIntervalMs:   5000,                      // 单位时间为5秒
				Threshold:        10,                        // 单位时间内容错数量
			},
		)
		sent.LoadRules()
		return new(UserClient)
	})
}

type UserClient struct {
}

func (u *UserClient) InitUserTenant(ctx context.Context, dto *client.UserTenantInitCmd) error {
	_, err := protoactor.ClientChain(
		protoactor.ClientActionArgs{
			Ctx:           ctx,
			OperationName: "User",
			Peer:          "InitUserTenant",
			SetTraceId: func(key, value string) error {
				dto.Dto.TraceId = value
				return nil
			},
			TryFn: func() (interface{}, error) {
				callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
				grainClient := client.GetUserGrainClient(protoactor.Cluster, strconv.FormatUint(dto.TenantId, 10))
				action, err := grainClient.TenantInitAction(dto, callOpts)
				if err != nil {
					panic(err)
				}
				if !action.Rsp.Success {
					return nil, fmt.Errorf("[USER:" + action.Rsp.ErrCode + ":" + action.Rsp.ErrMessage + "]")
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

func (u *UserClient) CreateUserAction(ctx context.Context, dto *client.CreateUserCmd) error {
	_, err := protoactor.ClientChain(
		protoactor.ClientActionArgs{
			Ctx:           ctx,
			OperationName: "User",
			Peer:          "CreateUserAction",
			SetTraceId: func(key, value string) error {
				dto.Dto.TraceId = value
				return nil
			},
			TryFn: func() (interface{}, error) {
				callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
				grainClient := client.GetUserGrainClient(protoactor.Cluster, dto.AccKey)
				action, err := grainClient.CreateUserAction(dto, callOpts)
				if err != nil {
					panic(err)
				}
				if !action.Rsp.Success {
					return nil, fmt.Errorf("[USER:" + action.Rsp.ErrCode + ":" + action.Rsp.ErrMessage + "]")
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

func (u *UserClient) LoginAction(ctx context.Context, dto *client.UserLoginCmd) (*client.UserLoginData, error) {
	rs, err := protoactor.ClientChain(
		protoactor.ClientActionArgs{
			Ctx:           ctx,
			OperationName: "User",
			Peer:          "LoginAction",
			SetTraceId: func(key, value string) error {
				dto.Dto.TraceId = value
				return nil
			},
			TryFn: func() (interface{}, error) {
				callOpts := cluster.DefaultGrainCallOptions(protoactor.Cluster).WithTimeout(time.Second).WithRetry(1)
				grainClient := client.GetUserGrainClient(protoactor.Cluster, dto.AccKey)
				action, err := grainClient.LoginAction(dto, callOpts)
				if err != nil {
					panic(err)
				}
				if !action.Rsp.Success {
					return nil, fmt.Errorf("[USER:" + action.Rsp.ErrCode + ":" + action.Rsp.ErrMessage + "]")
				} else {
					return action.Data, nil
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
		return rs.(*client.UserLoginData), nil
	}
	return nil, err
}
