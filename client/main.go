package main

import (
    "net"
	"encoding/binary"
    "fmt"
    //"math/rand"
    "time"
    "zrpc/lib"
    "bytes"
    "io"
)


var lock bool

func Send(conn net.Conn, req lib.Request, timeout int, list []lib.Request) {
    bts, _ := req.MarshalMsg(nil)
    headSize := len(bts)
    head := make([]byte, 2)
    binary.BigEndian.PutUint16(head, uint16(headSize))
    _, _ = conn.Write(head)
    _, _ = conn.Write(bts)
    fmt.Println("send")
}

func Receive(conn net.Conn) {
    var (
		BUF_SIZE  = 65538
		HEAD_SIZE = 2
		buffer    = bytes.NewBuffer(make([]byte, 0, BUF_SIZE)) //buffer用来缓存读取到的数据
	)

	var res lib.Response

	for {
        fmt.Println("conn read........")
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

				_, err = res.UnmarshalMsg(body)

				if err != nil {
					fmt.Println(err)
				} else {
                    if res.ResponseStatusCode == lib.Suc {
                        lock = true
                        fmt.Println("Receive", res)
                        goto endfor
                    }
					//Call(conn, req.FuncName, req.RequestMap)
				}

				//fmt.Println(string(body))
			} else {
				//fmt.Println("break body", buffer.Len(), string(readBytes))
				break
			}
		}

		Ready = true
	}
    endfor:
}


func main() {
	conn, _ := net.Dial("tcp", ":1234")
    //defer conn.Close()
    list := make([]lib.Request, 2)
    for i:=0;i<2;i++ {
        lock = false 
        id := fmt.Sprintf("%d", time.Now().UnixNano())
        arg := "Test"

        request := map[string]string{"a":"b"}

        req := lib.Request{}
        req.Id = id
        req.FuncName = arg
        req.RequestMap = request
        fmt.Println(req)
        Send(conn, req, 10, list)

        /*timer1 := time.NewTimer(time.Second * time.Duration(3))
        for {
            if lock {
                break
            } else {
                <-timer1.C
                fmt.Println("next")
            }
        }*/
        Receive(conn)
    }

    //time.Sleep(10 * time.Second)
    /*timeUnix:=time.Now().Unix() 
    fmt.Println("I am is haha")
    for {
        if len(req.Response) < 1 {
            if time.Now().Unix() - timeUnix > 5 {
                fmt.Println("end...")
                break
            }
            continue
        }
        fmt.Println(req.Response)
        break
    }*/


}

