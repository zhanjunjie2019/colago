module github.com/zhanjunjie2019/colago/samples/test

go 1.16

replace (
	github.com/zhanjunjie2019/colago/common => ../../common
	github.com/zhanjunjie2019/colago/samples/auth-client => ../auth-client
	github.com/zhanjunjie2019/colago/samples/shared => ../shared
	github.com/zhanjunjie2019/colago/samples/user-client => ../user-client
)

require (
	github.com/SkyAPM/go2sky v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/zhanjunjie2019/colago/common v0.0.0-00010101000000-000000000000
	github.com/zhanjunjie2019/colago/samples/auth-client v0.0.0-00010101000000-000000000000
	github.com/zhanjunjie2019/colago/samples/shared v0.0.0-00010101000000-000000000000
	github.com/zhanjunjie2019/colago/samples/user-client v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f
	golang.org/x/sys v0.0.0-20210906170528-6f6e22806c34 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83 // indirect
	skywalking.apache.org/repo/goapi v0.0.0-20210820070710-e10b78bbf481 // indirect
)
