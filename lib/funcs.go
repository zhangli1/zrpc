package lib

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	glib "lib"
	"net"
	"reflect"
	"time"
)

type SendType int

const (
	RequestSendType SendType = iota
	ResponseSendType
)

//序列化+发送
func Send(conn net.Conn, req Request, res Response, sendType SendType) {
	magicNum := make([]byte, 4)
	binary.BigEndian.PutUint32(magicNum, 0x123456)

	var bts []byte
	//if RequestSendType == RequestSendType {
	bts, _ = req.MarshalMsg(nil)
	/*} else {
		bts, _ = res.MarshalMsg(nil)
	}*/
	headSize := len(bts)
	head := make([]byte, 2)
	binary.BigEndian.PutUint16(head, uint16(headSize))
	packetBuf := bytes.NewBuffer(magicNum)
	packetBuf.Write(head)
	packetBuf.Write(bts)
	_, _ = conn.Write(packetBuf.Bytes())
	fmt.Println("server send")
}

func (req Request) Test(conn net.Conn, res Response) interface{} {
	Log.Info(res)
	//res.Id = res.Id
	//res.FuncName = res.FuncName
	//res.ResponseMap = req.RequestMap
	req.RequestStatusCode = Suc
	Send(conn, req, res, ResponseSendType)
	fmt.Println("req status", res)
	return res
}

func packetSlitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 检查 atEOF 参数 和 数据包头部的四个字节是否 为 0x123456(我们定义的协议的魔数)
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == 0x123456 {
		var l int16
		// 读出 数据包中 实际数据 的长度(大小为 0 ~ 2^16)
		binary.Read(bytes.NewReader(data[4:6]), binary.BigEndian, &l)
		pl := int(l) + 6
		if pl <= len(data) {
			return pl, data[:pl], nil
		}
	}
	return
}

func Receive(conn net.Conn, req Request, res Response, sendType SendType) {
	result := bytes.NewBuffer(nil)
	var buf [65542]byte // 由于 标识数据包长度 的只有两个字节 故数据包最大为 2^16+4(魔数)+2(长度标识)

	for {
		fmt.Println("receive.............")
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
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
				//if sendType == RequestSendType {
				_, err = req.UnmarshalMsg(body)
				//} else {
				//	_, err = res.UnmarshalMsg(body)
				//}

				if err != nil {
					fmt.Println(err)
				} else {
					//if sendType == RequestSendType {
					Log.Info(fmt.Sprintf("%s %s %s %v %v", req.Id, "content", req.FuncName, req.RequestMap, req.RequestStatusCode))
					if sendType == ResponseSendType {
						Call(conn, req, res)
					}
					/*} else {
						Log.Info(fmt.Sprintf("%s %s %s %v %v", res.Id, "content", res.FuncName, res.ResponseMap, res.ResponseStatusCode))
					}*/
				}
			}
		}
		result.Reset()
	}
}

//定义控制器函数Map类型，便于后续快捷使用
type ControllerMapsType map[string]reflect.Value

//声明控制器函数Map类型变量
var ControllerMaps ControllerMapsType

func Call(conn net.Conn, req Request, res Response) {
	crMap := make(ControllerMapsType, 0)
	vf := reflect.ValueOf(&req)
	vft := vf.Type()

	mNum := vf.NumMethod()
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}

	parms := []reflect.Value{reflect.ValueOf(conn), reflect.ValueOf(res)}
	glib.Try(
		func() {
			crMap[req.FuncName].Call(parms)
		},
		func(e interface{}) {
			Log.Error("funcName", req.FuncName, " ERROR:", e)
		})
}
