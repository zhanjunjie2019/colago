module github.com/zhanjunjie2019/colago/samples/auth-domain

go 1.16

replace (
	github.com/zhanjunjie2019/colago/common => ../../common
	github.com/zhanjunjie2019/colago/samples/shared => ../shared
)

require (
	github.com/AsynkronIT/protoactor-go v0.0.0-20210819095145-4b4b14c686df
	github.com/google/uuid v1.3.0 // indirect
	github.com/zhanjunjie2019/colago/common v0.0.0-00010101000000-000000000000
	github.com/zhanjunjie2019/colago/samples/shared v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f // indirect
	golang.org/x/sys v0.0.0-20210906170528-6f6e22806c34 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83 // indirect
)
