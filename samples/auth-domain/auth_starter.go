package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/zhanjunjie2019/colago/common/conf"
	"github.com/zhanjunjie2019/colago/common/ioc"
	_ "github.com/zhanjunjie2019/colago/common/postgres"
	"github.com/zhanjunjie2019/colago/common/protoactor"
	"github.com/zhanjunjie2019/colago/common/sentinel"
	"github.com/zhanjunjie2019/colago/samples/auth-domain/app/executor"
	_ "github.com/zhanjunjie2019/colago/samples/auth-domain/infrastructure/gatewayimpl"
	"github.com/zhanjunjie2019/colago/samples/shared/client"
	"time"
)

func init() {
	conf.InitConfig("./config.json")
	ioc.BatchProvideFinal()
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

	err := ioc.GetContainer().Invoke(func(
		sentFilter *sentinel.SentinelFilter) {
		protoactor.InitClientFilters(
			sentFilter,
		)
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	select {}
}
