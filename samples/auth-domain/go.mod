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
)
