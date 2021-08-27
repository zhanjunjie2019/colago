package main

import (
	"e.coding.net/double-j/ego/colago/common/conf"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/app/executor"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

func init() {
	conf.InitConfig("./config.json")
	_ = ioc.InjectSimpleBeanFinal()
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

	select {}
}
