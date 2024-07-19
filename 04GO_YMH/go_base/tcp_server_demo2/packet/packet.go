package packet

import (
	"bytes"
	"fmt"
)

const (
	CommandConn   = iota + 0x01 // 0x01 连接请求包
	CommandSubmit               // 0x02 消息请求包
)

const (
	CommandAck       = iota + 0x81 // 0x81 连接请求的响应包
	CommandSubmitAck               // 0x82 连接请求的响应包
)

type Packet interface {
	Decode([]byte) error
	Encode() ([]byte, error)
}

type Submit struct {
	ID      string
	PayLoad []byte
}

func (s *Submit) Decode(pktBody []byte) error {
	s.ID = string(pktBody[:8])
	s.PayLoad = pktBody[8:]
	return nil
}

func (s *Submit) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), s.PayLoad}, nil), nil
}

type SubmitAck struct {
	ID     string
	Result uint8
}

func (s *SubmitAck) Decode(pktBody []byte) error {
	s.ID = string(pktBody[:8])
	s.Result = uint8(pktBody[8])
	return nil
}

func (s *SubmitAck) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), {s.Result}}, nil), nil
}

// 这里各种类的编解码被调用的前提，是明确数据流是什么类型的，因此我们需要在
// 包级提供一个导出函数Decode，这个函数负责从字节流中解析出来对应的类型（根据commandId）并调用Decode方法

func Decode(packet []byte) (Packet, error) {
	commandID := packet[0]
	pktBody := packet[1:]
	switch commandID {
	case CommandConn:
		return nil, nil
	case CommandAck:
		return nil, nil
	case CommandSubmit:
		s := Submit{}
		err := s.Decode(pktBody)
		if err != nil {
			return nil, err
		}
		return &s, nil
	case CommandSubmitAck:
		s := SubmitAck{}
		err := s.Decode(pktBody)
		if err != nil {
			return nil, err
		}
		return &s, nil
	default:
		return nil, fmt.Errorf("unknown commandID [%d]", commandID)
	}
}

// 同时，我们也需要包级的Encode函数，根据传入的packet类型调用对应的Encode方法实现对象的编码:

func Encode(p Packet) ([]byte, error) {
	var commandId uint8
	var pktBody []byte
	var err error
	switch t := p.(type) {
	case *Submit:
		commandId = CommandSubmit
		pktBody, err = p.Encode()
		if err != nil {
			return nil, err
		}
	case *SubmitAck:
		commandId = CommandSubmitAck
		pktBody, err = p.Encode()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unkonw type [%s]", t)
	}
	return bytes.Join([][]byte{{commandId}, pktBody}, nil), nil
}
