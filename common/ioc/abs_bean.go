package ioc

type AbsBean interface {
	New() AbsBean
}
