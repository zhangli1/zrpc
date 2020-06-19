package lib

//go:generate /Users/zhangli/work/golang/work/hqs_monitor/bin/msgp

type State int

const (
	Wait State = iota
	Fail
	Suc
)

type Request struct {
	Id                string
	FuncName          string
	RequestMap        map[string]string
	RequestStatusCode State
}

type Response struct {
	Id                 string
	FuncName           string
	ResponseMap        interface{}
	ResponseStatusCode State
}
