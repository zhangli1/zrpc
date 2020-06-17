package lib

//go:generate /Users/zhangli/work/golang/work/hqs_monitor/bin/msgp

type Foo struct {
	FuncName string
	Request  map[string]string
	Response map[string]interface{}
}
