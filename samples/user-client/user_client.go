package userclient

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/common/sentinuel"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"strconv"
	"time"
)

func init() {
	err := ioc.InjectSimpleBean(new(UserClient))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type UserClient struct {
	sent *sentinuel.Sentinel `ij:"sentinuel.Sentinel"`
}

func (u *UserClient) New() ioc.AbsBean {
	u.sent.AppendCircuitbreakerRules(&circuitbreaker.Rule{
		Resource:         "User.InitAuthTenant",
		Strategy:         circuitbreaker.ErrorCount, // 异常记数方案
		RetryTimeoutMs:   3000,                      // 熔断后3秒重试
		MinRequestAmount: 10,                        // 单位时间内10个请求以上才进入异常记数计算
		StatIntervalMs:   5000,                      // 单位时间为5秒
		Threshold:        10,                        // 单位时间内容错数量
	})
	u.sent.AppendCircuitbreakerRules(&circuitbreaker.Rule{
		Resource:         "User.CreateUserAction",
		Strategy:         circuitbreaker.ErrorCount, // 异常记数方案
		RetryTimeoutMs:   3000,                      // 熔断后3秒重试
		MinRequestAmount: 10,                        // 单位时间内10个请求以上才进入异常记数计算
		StatIntervalMs:   5000,                      // 单位时间为5秒
		Threshold:        10,                        // 单位时间内容错数量
	})
	u.sent.AppendCircuitbreakerRules(&circuitbreaker.Rule{
		Resource:         "User.LoginAction",
		Strategy:         circuitbreaker.ErrorCount, // 异常记数方案
		RetryTimeoutMs:   3000,                      // 熔断后3秒重试
		MinRequestAmount: 10,                        // 单位时间内10个请求以上才进入异常记数计算
		StatIntervalMs:   5000,                      // 单位时间为5秒
		Threshold:        10,                        // 单位时间内容错数量
	})
	u.sent.LoadRules()
	return u
}

func (u *UserClient) InitUserTenant(dto *client.UserTenantInitCmd) error {
	_, err := u.sent.Entry(
		"User.InitAuthTenant",
		func() (interface{}, error) {
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
		func(er interface{}) (interface{}, error) {
			er2, ok := er.(error)
			if ok {
				return nil, er2
			}
			return nil, fmt.Errorf("%v", er)
		},
	)
	return err
}

func (u *UserClient) CreateUserAction(dto *client.CreateUserCmd) error {
	_, err := u.sent.Entry(
		"User.CreateUserAction",
		func() (interface{}, error) {
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
		func(er interface{}) (interface{}, error) {
			er2, ok := er.(error)
			if ok {
				return nil, er2
			}
			return nil, fmt.Errorf("%v", er)
		},
	)
	return err
}

func (u *UserClient) LoginAction(dto *client.UserLoginCmd) (*client.UserLoginData, error) {
	rs, err := u.sent.Entry(
		"User.LoginAction",
		func() (interface{}, error) {
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
		func(er interface{}) (interface{}, error) {
			er2, ok := er.(error)
			if ok {
				return nil, er2
			}
			return nil, fmt.Errorf("%v", er)
		},
	)
	if rs != nil {
		return rs.(*client.UserLoginData), nil
	}
	return nil, err
}
