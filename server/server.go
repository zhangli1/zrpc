package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"reflect"
	"time"
	"zrpc/lib"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func doConn(conn net.Conn) {
	var (
		BUF_SIZE  = 65538
		HEAD_SIZE = 2
		buffer    = bytes.NewBuffer(make([]byte, 0, BUF_SIZE)) //buffer用来缓存读取到的数据
	)

	var foo lib.Foo

	for {
		readBytes := make([]byte, BUF_SIZE)
		Ready := false
		//首先读取数据
		readByteNum, err := conn.Read(readBytes)
		if err != nil {
			if err == io.EOF {
				//fmt.Println("EOF")
				time.Sleep(1 * time.Second)
				continue
			} else {
				if Ready {
					time.Sleep(1 * time.Second)
					break
				}
			}
		}
		buffer.Write(readBytes[0:readByteNum]) //将读取到的数据放到buffer中
		//fmt.Println("readBytes", string(readBytes))

		for {
			bodyLen := 0
			head := make([]byte, HEAD_SIZE)
			_, _ = buffer.Read(head)

			bodyLen = int(binary.BigEndian.Uint16(head))
			fmt.Println("bodyLen", bodyLen)

			if buffer.Len() >= bodyLen && bodyLen > 0 {
				body := make([]byte, bodyLen)
				_, _ = buffer.Read(body)

				_, err = foo.UnmarshalMsg(body)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("content", foo.FuncName, foo.Request, foo.Response)
					Log.Info(fmt.Sprintf("%s %s %v %v", "content", foo.FuncName, foo.Request, foo.Response))

					Call(conn, foo.FuncName, foo.Request)
				}

				//fmt.Println(string(body))
			} else {
				//fmt.Println("break body", buffer.Len())
				break
			}
		}

		Ready = true

	}
}

//定义控制器函数Map类型，便于后续快捷使用
type ControllerMapsType map[string]reflect.Value

//声明控制器函数Map类型变量
var ControllerMaps ControllerMapsType

func Call(conn net.Conn, method string, request map[string]string) {
	var foo lib.Foo
	crMap := make(ControllerMapsType, 0)
	vf := reflect.ValueOf(&foo)
	vft := vf.Type()

	mNum := vf.NumMethod()
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}

	testStr := request
	parms := []reflect.Value{reflect.ValueOf(conn), reflect.ValueOf(testStr)}
	crMap[method].Call(parms)

}

func init() {
	path, _ := os.Getwd()
	lib.ConfigLocalFilesystemLogger(Log, path, "log/go.log", 24*30*time.Hour, 1*time.Hour)
	Log.Info("iaaaaa")
}

func startup(host string, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		Log.Fatal(err)
		return
	}
	Log.Info("start listening on 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			Log.Fatal(err)
			return
		}
		go doConn(conn)
	}
}
