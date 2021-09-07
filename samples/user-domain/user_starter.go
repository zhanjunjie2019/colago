package main

import (
	"e.coding.net/double-j/ego/colago/common/conf"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/common/sentinel"
	"e.coding.net/double-j/ego/colago/common/skywalking"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/app/executor"
	_ "e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/gatewayimpl"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

func init() {
	conf.InitConfig("./config.json")
	err := ioc.InjectSimpleBeanFinal()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func main() {
	client.UserFactory(executor.NewUserAppExe)

	protoactor.InitConsulActor(
		"127.0.0.1:8500",
		"colago-samples",
		0,
		map[string]actor.Actor{
			"User": &client.UserActor{
				Timeout: time.Second,
			},
		},
	)

	err := skywalking.NewGlobalTracer("user-service", "127.0.0.1:11800")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	protoactor.InitClientFilters(
		sentinel.SentinulFilterFactory,
		skywalking.SkyFilterFactory,
	)

	select {}
}
