// Package client is generated by protoactor-go/protoc-gen-gograin@0.1.0
package client

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	logmod "github.com/AsynkronIT/protoactor-go/log"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/gogo/protobuf/proto"
)

var (
	plog = logmod.New(logmod.InfoLevel, "[GRAIN]")
	_    = proto.Marshal
	_    = fmt.Errorf
	_    = math.Inf
)

// SetLogLevel sets the log level.
func SetLogLevel(level logmod.Level) {
	plog.SetLevel(level)
}

var xUserFactory func() User

// UserFactory produces a User
func UserFactory(factory func() User) {
	xUserFactory = factory
}

// GetUserGrainClient instantiates a new UserGrainClient with given ID
func GetUserGrainClient(c *cluster.Cluster, id string) *UserGrainClient {
	if c == nil {
		panic(fmt.Errorf("nil cluster instance"))
	}
	if id == "" {
		panic(fmt.Errorf("empty id"))
	}
	return &UserGrainClient{ID: id, cluster: c}
}

// User interfaces the services available to the User
type User interface {
	Init(id string)
	Terminate()
	ReceiveDefault(ctx actor.Context)
	LoginAction(*UserLoginCmd, cluster.GrainContext) (*UserLoginResponse, error)
}

// UserGrainClient holds the base data for the UserGrain
type UserGrainClient struct {
	ID      string
	cluster *cluster.Cluster
}

// LoginAction requests the execution on to the cluster with CallOptions
func (g *UserGrainClient) LoginAction(r *UserLoginCmd, opts ...*cluster.GrainCallOptions) (*UserLoginResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 0, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "User", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &UserLoginResponse{}
		err = proto.Unmarshal(msg.MessageData, result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case *cluster.GrainErrorResponse:
		if msg.Code == remote.ResponseStatusCodeDeadLetter.ToInt32() {
			return nil, remote.ErrDeadLetter
		}
		return nil, errors.New(msg.Err)
	default:
		return nil, errors.New("unknown response")
	}
}

// UserActor represents the actor structure
type UserActor struct {
	inner   User
	Timeout time.Duration
}

// Receive ensures the lifecycle of the actor for the received message
func (a *UserActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
	case *cluster.ClusterInit:
		a.inner = xUserFactory()
		a.inner.Init(msg.ID)
		if a.Timeout > 0 {
			ctx.SetReceiveTimeout(a.Timeout)
		}

	case *actor.ReceiveTimeout:
		a.inner.Terminate()
		ctx.Poison(ctx.Self())

	case actor.AutoReceiveMessage: // pass
	case actor.SystemMessage: // pass

	case *cluster.GrainRequest:
		switch msg.MethodIndex {
		case 0:
			req := &UserLoginCmd{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				plog.Error("LoginAction(UserLoginCmd) proto.Unmarshal failed.", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			r0, err := a.inner.LoginAction(req, ctx)
			if err != nil {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			bytes, err := proto.Marshal(r0)
			if err != nil {
				plog.Error("LoginAction(UserLoginCmd) proto.Marshal failed", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			resp := &cluster.GrainResponse{MessageData: bytes}
			ctx.Respond(resp)

		}
	default:
		a.inner.ReceiveDefault(ctx)
	}
}
