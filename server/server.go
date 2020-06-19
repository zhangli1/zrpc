package server

import (
	"fmt"
	"net"
	"os"
	"time"
	"zrpc/lib"
)

func handleConn2(conn net.Conn) {
	var req lib.Request
	var res lib.Response
	lib.Receive(conn, req, res, lib.ServerSendType)
}

func init() {
	path, _ := os.Getwd()
	lib.Log.SetReportCaller(true)
	lib.ConfigLocalFilesystemLogger(lib.Log, path, "log/go.log", 24*30*time.Hour, 1*time.Hour)
}

func startup(host string, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		lib.Log.Fatal(err)
		return
	}
	lib.Log.Info("start listening on ", host, ":", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			lib.Log.Fatal(err)
			return
		}
		//go doConn(conn)
		go handleConn2(conn)
	}
}
