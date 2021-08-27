package protoactor

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/AsynkronIT/protoactor-go/cluster/consul"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/hashicorp/consul/api"
)

var (
	system  = actor.NewActorSystem()
	Cluster *cluster.Cluster
)

func InitConsulActor(consulHost string, clusterName string, serverPort int, kindMaps map[string]actor.Actor) {
	provider, _ := consul.NewWithConfig(&api.Config{
		Address: consulHost,
	})

	kinds := make([]*cluster.Kind, 0)
	for k, v := range kindMaps {
		kind := cluster.NewKind(k, actor.PropsFromProducer(func() actor.Actor {
			return v
		}))
		kinds = append(kinds, kind)
	}

	remoteCfg := remote.Configure("127.0.0.1", serverPort)
	configure := cluster.Configure(clusterName, provider, remoteCfg, kinds...)
	Cluster = cluster.New(system, configure)
	Cluster.Start()
}
