package lib

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Request) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "FuncName":
			z.FuncName, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "FuncName")
				return
			}
		case "RequestMap":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "RequestMap")
				return
			}
			if z.RequestMap == nil {
				z.RequestMap = make(map[string]string, zb0002)
			} else if len(z.RequestMap) > 0 {
				for key := range z.RequestMap {
					delete(z.RequestMap, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 string
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "RequestMap")
					return
				}
				za0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "RequestMap", za0001)
					return
				}
				z.RequestMap[za0001] = za0002
			}
		case "RequestStatusCode":
			{
				var zb0003 int
				zb0003, err = dc.ReadInt()
				if err != nil {
					err = msgp.WrapError(err, "RequestStatusCode")
					return
				}
				z.RequestStatusCode = State(zb0003)
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Request) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Id"
	err = en.Append(0x84, 0xa2, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	// write "FuncName"
	err = en.Append(0xa8, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.FuncName)
	if err != nil {
		err = msgp.WrapError(err, "FuncName")
		return
	}
	// write "RequestMap"
	err = en.Append(0xaa, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.RequestMap)))
	if err != nil {
		err = msgp.WrapError(err, "RequestMap")
		return
	}
	for za0001, za0002 := range z.RequestMap {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "RequestMap")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "RequestMap", za0001)
			return
		}
	}
	// write "RequestStatusCode"
	err = en.Append(0xb1, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(int(z.RequestStatusCode))
	if err != nil {
		err = msgp.WrapError(err, "RequestStatusCode")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Request) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Id"
	o = append(o, 0x84, 0xa2, 0x49, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "FuncName"
	o = append(o, 0xa8, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FuncName)
	// string "RequestMap"
	o = append(o, 0xaa, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.RequestMap)))
	for za0001, za0002 := range z.RequestMap {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	// string "RequestStatusCode"
	o = append(o, 0xb1, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt(o, int(z.RequestStatusCode))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Request) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "FuncName":
			z.FuncName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FuncName")
				return
			}
		case "RequestMap":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RequestMap")
				return
			}
			if z.RequestMap == nil {
				z.RequestMap = make(map[string]string, zb0002)
			} else if len(z.RequestMap) > 0 {
				for key := range z.RequestMap {
					delete(z.RequestMap, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RequestMap")
					return
				}
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RequestMap", za0001)
					return
				}
				z.RequestMap[za0001] = za0002
			}
		case "RequestStatusCode":
			{
				var zb0003 int
				zb0003, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RequestStatusCode")
					return
				}
				z.RequestStatusCode = State(zb0003)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Request) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.Id) + 9 + msgp.StringPrefixSize + len(z.FuncName) + 11 + msgp.MapHeaderSize
	if z.RequestMap != nil {
		for za0001, za0002 := range z.RequestMap {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += 18 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Response) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "FuncName":
			z.FuncName, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "FuncName")
				return
			}
		case "ResponseMap":
			z.ResponseMap, err = dc.ReadIntf()
			if err != nil {
				err = msgp.WrapError(err, "ResponseMap")
				return
			}
		case "ResponseStatusCode":
			{
				var zb0002 int
				zb0002, err = dc.ReadInt()
				if err != nil {
					err = msgp.WrapError(err, "ResponseStatusCode")
					return
				}
				z.ResponseStatusCode = State(zb0002)
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Response) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Id"
	err = en.Append(0x84, 0xa2, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	// write "FuncName"
	err = en.Append(0xa8, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.FuncName)
	if err != nil {
		err = msgp.WrapError(err, "FuncName")
		return
	}
	// write "ResponseMap"
	err = en.Append(0xab, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x70)
	if err != nil {
		return
	}
	err = en.WriteIntf(z.ResponseMap)
	if err != nil {
		err = msgp.WrapError(err, "ResponseMap")
		return
	}
	// write "ResponseStatusCode"
	err = en.Append(0xb2, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(int(z.ResponseStatusCode))
	if err != nil {
		err = msgp.WrapError(err, "ResponseStatusCode")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Response) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Id"
	o = append(o, 0x84, 0xa2, 0x49, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "FuncName"
	o = append(o, 0xa8, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FuncName)
	// string "ResponseMap"
	o = append(o, 0xab, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x70)
	o, err = msgp.AppendIntf(o, z.ResponseMap)
	if err != nil {
		err = msgp.WrapError(err, "ResponseMap")
		return
	}
	// string "ResponseStatusCode"
	o = append(o, 0xb2, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt(o, int(z.ResponseStatusCode))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Response) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "FuncName":
			z.FuncName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FuncName")
				return
			}
		case "ResponseMap":
			z.ResponseMap, bts, err = msgp.ReadIntfBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ResponseMap")
				return
			}
		case "ResponseStatusCode":
			{
				var zb0002 int
				zb0002, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "ResponseStatusCode")
					return
				}
				z.ResponseStatusCode = State(zb0002)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Response) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.Id) + 9 + msgp.StringPrefixSize + len(z.FuncName) + 12 + msgp.GuessSize(z.ResponseMap) + 19 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *State) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = State(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z State) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt(int(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z State) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *State) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = State(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z State) Msgsize() (s int) {
	s = msgp.IntSize
	return
}
