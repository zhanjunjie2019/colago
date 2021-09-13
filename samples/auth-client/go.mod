module github.com/zhanjunjie2019/colago/samples/auth-client

go 1.16

replace (
	github.com/zhanjunjie2019/colago/common => ../../common
	github.com/zhanjunjie2019/colago/samples/shared => ../shared
)

require (
	github.com/AsynkronIT/protoactor-go v0.0.0-20210819095145-4b4b14c686df
	github.com/alibaba/sentinel-golang v1.0.3
	github.com/zhanjunjie2019/colago/common v0.0.0-00010101000000-000000000000
	github.com/zhanjunjie2019/colago/samples/shared v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
)
