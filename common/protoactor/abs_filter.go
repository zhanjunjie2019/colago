package protoactor

type ActorClientFilter interface {
	SetNext(filter ActorClientFilter)
	Filter(clientActionArgs ClientActionArgs) (rs interface{}, err error)
}
