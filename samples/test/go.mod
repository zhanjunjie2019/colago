module e.coding.net/double-j/ego/colago/samples/test

go 1.16

replace (
	e.coding.net/double-j/ego/colago/common => ../../common
	e.coding.net/double-j/ego/colago/samples/auth-client => ../auth-client
	e.coding.net/double-j/ego/colago/samples/shared => ../shared
	e.coding.net/double-j/ego/colago/samples/user-client => ../user-client
)

require (
	e.coding.net/double-j/ego/colago/common v0.0.0-00010101000000-000000000000
	e.coding.net/double-j/ego/colago/samples/auth-client v0.0.0-00010101000000-000000000000
	e.coding.net/double-j/ego/colago/samples/shared v0.0.0-00010101000000-000000000000
	e.coding.net/double-j/ego/colago/samples/user-client v0.0.0-00010101000000-000000000000
)
