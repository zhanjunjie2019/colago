package protoactor

var rootClientFilter = new(RootActorClientFilter)
var finalClientFilter = new(FinalActorClientFilter)

func InitClientFilters(filterFactorys ...func() ActorClientFilter) {
	if filterFactorys != nil {
		var lastFilter ActorClientFilter
		for _, filterFactory := range filterFactorys {
			nextFilter := filterFactory()
			if lastFilter == nil {
				rootClientFilter.SetNext(nextFilter)
			} else {
				lastFilter.SetNext(nextFilter)
			}
			lastFilter = nextFilter
		}
		lastFilter.SetNext(finalClientFilter)
	} else {
		rootClientFilter.SetNext(finalClientFilter)
	}
}

func ClientChain(clientDing ClientActionArgs) (rs interface{}, err error) {
	if rootClientFilter.next == nil {
		rootClientFilter.SetNext(finalClientFilter)
	}
	return rootClientFilter.Filter(clientDing)
}
