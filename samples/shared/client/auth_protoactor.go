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
	//plog = logmod.New(logmod.InfoLevel, "[GRAIN]")
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// SetLogLevel sets the log level.
//func SetLogLevel(level logmod.Level) {
//	plog.SetLevel(level)
//}

var xAuthFactory func() Auth

// AuthFactory produces a Auth
func AuthFactory(factory func() Auth) {
	xAuthFactory = factory
}

// GetAuthGrainClient instantiates a new AuthGrainClient with given ID
func GetAuthGrainClient(c *cluster.Cluster, id string) *AuthGrainClient {
	if c == nil {
		panic(fmt.Errorf("nil cluster instance"))
	}
	if id == "" {
		panic(fmt.Errorf("empty id"))
	}
	return &AuthGrainClient{ID: id, cluster: c}
}

// Auth interfaces the services available to the Auth
type Auth interface {
	Init(id string)
	Terminate()
	ReceiveDefault(ctx actor.Context)
	TenantInitAction(*AuthTenantInitCmd, cluster.GrainContext) (*AuthResponse, error)
	CreateAuthAction(*CreateAuthCmd, cluster.GrainContext) (*AuthResponse, error)
	FindRolesByUserId(*RoleQry, cluster.GrainContext) (*RoleQryResponse, error)
	FindAuthsByUserId(*AuthQry, cluster.GrainContext) (*AuthQryResponse, error)
}

// AuthGrainClient holds the base data for the AuthGrain
type AuthGrainClient struct {
	ID      string
	cluster *cluster.Cluster
}

// TenantInitAction requests the execution on to the cluster with CallOptions
func (g *AuthGrainClient) TenantInitAction(r *AuthTenantInitCmd, opts ...*cluster.GrainCallOptions) (*AuthResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 0, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Auth", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &AuthResponse{}
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

// CreateAuthAction requests the execution on to the cluster with CallOptions
func (g *AuthGrainClient) CreateAuthAction(r *CreateAuthCmd, opts ...*cluster.GrainCallOptions) (*AuthResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 1, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Auth", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &AuthResponse{}
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

// FindRolesByUserId requests the execution on to the cluster with CallOptions
func (g *AuthGrainClient) FindRolesByUserId(r *RoleQry, opts ...*cluster.GrainCallOptions) (*RoleQryResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 2, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Auth", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &RoleQryResponse{}
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

// FindAuthsByUserId requests the execution on to the cluster with CallOptions
func (g *AuthGrainClient) FindAuthsByUserId(r *AuthQry, opts ...*cluster.GrainCallOptions) (*AuthQryResponse, error) {
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqMsg := &cluster.GrainRequest{MethodIndex: 3, MessageData: bytes}
	resp, err := g.cluster.Call(g.ID, "Auth", reqMsg, opts...)
	if err != nil {
		return nil, err
	}
	switch msg := resp.(type) {
	case *cluster.GrainResponse:
		result := &AuthQryResponse{}
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

// AuthActor represents the actor structure
type AuthActor struct {
	inner   Auth
	Timeout time.Duration
}

// Receive ensures the lifecycle of the actor for the received message
func (a *AuthActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
	case *cluster.ClusterInit:
		a.inner = xAuthFactory()
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
			req := &AuthTenantInitCmd{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				plog.Error("TenantInitAction(AuthTenantInitCmd) proto.Unmarshal failed.", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			r0, err := a.inner.TenantInitAction(req, ctx)
			if err != nil {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			bytes, err := proto.Marshal(r0)
			if err != nil {
				plog.Error("TenantInitAction(AuthTenantInitCmd) proto.Marshal failed", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			resp := &cluster.GrainResponse{MessageData: bytes}
			ctx.Respond(resp)
		case 1:
			req := &CreateAuthCmd{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				plog.Error("CreateAuthAction(CreateAuthCmd) proto.Unmarshal failed.", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			r0, err := a.inner.CreateAuthAction(req, ctx)
			if err != nil {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			bytes, err := proto.Marshal(r0)
			if err != nil {
				plog.Error("CreateAuthAction(CreateAuthCmd) proto.Marshal failed", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			resp := &cluster.GrainResponse{MessageData: bytes}
			ctx.Respond(resp)
		case 2:
			req := &RoleQry{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				plog.Error("FindRolesByUserId(RoleQry) proto.Unmarshal failed.", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			r0, err := a.inner.FindRolesByUserId(req, ctx)
			if err != nil {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			bytes, err := proto.Marshal(r0)
			if err != nil {
				plog.Error("FindRolesByUserId(RoleQry) proto.Marshal failed", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			resp := &cluster.GrainResponse{MessageData: bytes}
			ctx.Respond(resp)
		case 3:
			req := &AuthQry{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				plog.Error("FindAuthsByUserId(AuthQry) proto.Unmarshal failed.", logmod.Error(err))
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			r0, err := a.inner.FindAuthsByUserId(req, ctx)
			if err != nil {
				resp := &cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
				return
			}
			bytes, err := proto.Marshal(r0)
			if err != nil {
				plog.Error("FindAuthsByUserId(AuthQry) proto.Marshal failed", logmod.Error(err))
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
