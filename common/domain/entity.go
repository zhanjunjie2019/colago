package domain

type Entity struct {
	id uint64
}

func (e *Entity) Id() uint64 {
	return e.id
}

func (e *Entity) SetId(id uint64) {
	e.id = id
}
