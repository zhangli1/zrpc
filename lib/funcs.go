package lib

import (
	"encoding/binary"
	"fmt"
	"net"
)

//序列化+发送
func (res Response) Send(conn net.Conn) bool {
	bts, err := res.MarshalMsg(nil)
	if err != nil {
		Log.Error(err)
		return false
	}
	headSize := len(bts)
	head := make([]byte, 2)
	binary.BigEndian.PutUint16(head, uint16(headSize))
	_, err = conn.Write(head)
	if err != nil {
		Log.Error(err)
		return false
	}
	_, err = conn.Write(bts)
	if err != nil {
		Log.Error(err)
		return false
	}
	return true
}

func (req Request) Test(conn net.Conn, r Request) interface{} {
	req.RequestStatusCode = Suc
	fmt.Println("test test")
	Log.Info(r)

	res := Response{}
	res.Id = r.Id
	res.FuncName = r.FuncName
	res.ResponseMap = r.RequestMap
	res.ResponseStatusCode = Suc
	/*bts, _ := foo.MarshalMsg(nil)
	str := glib.MapToJson(request)
	_, _ = conn.Write([]byte(str))
	*/
	fmt.Println("req status", res, res.Send(conn))
	return r
}
