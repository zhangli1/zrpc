package lib

import (
	glib "lib"
	"net"
)

func (foo Foo) Test(conn net.Conn, request map[string]string) interface{} {
	//Log.Info(request)
	str := glib.MapToJson(request)
	_, _ = conn.Write([]byte(str))
	return request
}
