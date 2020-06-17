package lib

import (
	"fmt"
	"net"
)

func (foo Foo) Test(conn net.Conn, request map[string]string) interface{} {
	fmt.Println("hahaha", request)
	_, _ = conn.Write([]byte("hahaha"))
	return request
}
