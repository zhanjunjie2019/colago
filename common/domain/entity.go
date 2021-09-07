package domain

import "golang.org/x/net/context"

type Entity struct {
	id  uint64
	ctx context.Context
}

func (e *Entity) Id() uint64 {
	return e.id
}

func (e *Entity) SetId(id uint64) {
	e.id = id
}

func (e *Entity) Ctx() context.Context {
	return e.ctx
}

func (e *Entity) SetCtx(ctx context.Context) {
	e.ctx = ctx
}
