package main

import (
	"e.coding.net/double-j/ego/colago/common/conf"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/app/executor"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

func init() {
	conf.InitConfig("./config.json")
	_ = ioc.InjectSimpleBeanFinal()
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

	select {}
}
