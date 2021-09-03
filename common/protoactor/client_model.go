package protoactor

type ClientActionArgs struct {
	Resource string
	TryFn    func() (interface{}, error)
	CatchFn  func(interface{}) (interface{}, error)
}
