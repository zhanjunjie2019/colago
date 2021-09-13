module e.coding.net/double-j/ego/colago/samples/auth-client

go 1.16

replace (
	e.coding.net/double-j/ego/colago/common => ../../common
	e.coding.net/double-j/ego/colago/samples/shared => ../shared
)

require (
	e.coding.net/double-j/ego/colago/common v0.0.0-00010101000000-000000000000
	e.coding.net/double-j/ego/colago/samples/shared v0.0.0-00010101000000-000000000000
	github.com/AsynkronIT/protoactor-go v0.0.0-20210819095145-4b4b14c686df
	github.com/alibaba/sentinel-golang v1.0.3
	go.uber.org/dig v1.12.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
)
