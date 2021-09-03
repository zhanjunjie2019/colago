package protoactor

type RootActorClientFilter struct {
	next ActorClientFilter
}

func (r *RootActorClientFilter) SetNext(filter ActorClientFilter) {
	r.next = filter
}

func (r *RootActorClientFilter) Filter(clientActionArgs ClientActionArgs) (rs interface{}, err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			rs, err = clientActionArgs.CatchFn(recoverErr)
		}
	}()
	return r.next.Filter(clientActionArgs)
}
