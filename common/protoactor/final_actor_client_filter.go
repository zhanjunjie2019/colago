package protoactor

type FinalActorClientFilter struct {
}

func (f *FinalActorClientFilter) SetNext(filter ActorClientFilter) {
	panic("FinalActorClientFilter cannot has next filter")
}

func (f *FinalActorClientFilter) Filter(clientActionArgs ClientActionArgs) (rs interface{}, err error) {
	return clientActionArgs.TryFn()
}
