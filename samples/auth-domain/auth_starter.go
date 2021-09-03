package main

import (
	"e.coding.net/double-j/ego/colago/common/conf"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/common/sentinel"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/app/executor"
	_ "e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/gatewayimpl"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
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
	client.AuthFactory(executor.NewAuthAppExe)

	protoactor.InitConsulActor(
		"127.0.0.1:8500",
		"colago-samples",
		0,
		map[string]actor.Actor{
			"Auth": &client.AuthActor{
				Timeout: time.Second,
			},
		},
	)

	protoactor.InitClientFilters(sentinel.SentinuelActorChainFactory)

	select {}
}
