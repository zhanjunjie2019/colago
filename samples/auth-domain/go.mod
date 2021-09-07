module e.coding.net/double-j/ego/colago/samples/auth-domain

go 1.16

replace (
	e.coding.net/double-j/ego/colago/common => ../../common
	e.coding.net/double-j/ego/colago/samples/shared => ../shared
)

require (
	e.coding.net/double-j/ego/colago/common v0.0.0-00010101000000-000000000000
	e.coding.net/double-j/ego/colago/samples/shared v0.0.0-00010101000000-000000000000
	github.com/AsynkronIT/protoactor-go v0.0.0-20210819095145-4b4b14c686df
	github.com/SkyAPM/go2sky v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f // indirect
	golang.org/x/sys v0.0.0-20210906170528-6f6e22806c34 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83 // indirect
	skywalking.apache.org/repo/goapi v0.0.0-20210820070710-e10b78bbf481 // indirect
)
