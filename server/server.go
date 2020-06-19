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
	lib.Receive(conn, req, res, lib.ResponseSendType)
	//defer conn.Close()
	//defer fmt.Println("关闭")
	//fmt.Println("新连接：", conn.RemoteAddr())
	/*result := bytes.NewBuffer(nil)
	var buf [65542]byte // 由于 标识数据包长度 的只有两个字节 故数据包最大为 2^16+4(魔数)+2(长度标识)

	var req lib.Request
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		} else {
			scanner := bufio.NewScanner(result)
			scanner.Split(packetSlitFunc)
			for scanner.Scan() {
				fmt.Println("recv:", string(scanner.Bytes()[6:]))
				body := scanner.Bytes()[6:]
				_, err = req.UnmarshalMsg(body)

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(req.Id, "content2", req.FuncName, req.RequestMap)
					lib.Log.Info(fmt.Sprintf("%s %s %s %v %v", req.Id, "content", req.FuncName, req.RequestMap))

					Call(conn, req)
				}
			}
		}
		result.Reset()
	}*/
}

/*func doConn(conn net.Conn) {
	var (
		BUF_SIZE  = 65538
		HEAD_SIZE = 2
		buffer    = bytes.NewBuffer(make([]byte, 0, BUF_SIZE)) //buffer用来缓存读取到的数据
	)

	var req lib.Request

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
			fmt.Println("buffer.Len()", buffer.Len(), "bodyLen", bodyLen)

			if buffer.Len() >= bodyLen && bodyLen > 0 {
				body := make([]byte, bodyLen)
				_, _ = buffer.Read(body)

				_, err = req.UnmarshalMsg(body)

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(req.Id, "content2", req.FuncName, req.RequestMap)
					lib.Log.Info(fmt.Sprintf("%s %s %s %v %v", req.Id, "content", req.FuncName, req.RequestMap))

					Call(conn, req)
				}

				//fmt.Println(string(body))
			} else {
				fmt.Println("break body", buffer.Len(), string(readBytes))
				buffer.Reset()
				break
			}
		}

		Ready = true

	}
}*/

func init() {
	path, _ := os.Getwd()
	//lib.Log.SetOutput(lib.Log.Writer())
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
