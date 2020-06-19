package main

import (
    "net"
    "fmt"
    "time"
    "zrpc/lib"
)

var MaxNum int = 50000

func main() {
	conn, _ := net.Dial("tcp", ":1234")
    //defer conn.Close()
    // 发送超时
    //conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
    // 读取超时
    //conn.SetReadDeadline(time.Now().Add(5 * time.Second))

    var req lib.Request
    var res lib.Response
    for i:=0;i<MaxNum;i++ {
        func() {
            id := fmt.Sprintf("%d", time.Now().UnixNano())
            arg := "Test"

            request := map[string]string{"a":"b"}

            req.Id = id
            req.FuncName = arg
            req.RequestMap = request
            fmt.Println(req)

            lib.Send(conn, req, res, lib.ClientSendType)
            //lib.List[req.Id] = req
        }()
    }
    list := make(map[string]lib.Request, MaxNum)
    go lib.Receive(conn, req, res, lib.ClientSendType, list)

    for {
        fmt.Println(list)
        if len(list) >= MaxNum {
            fmt.Println("end...")
            break
        }
        time.Sleep(1 * time.Second)
    }
}

