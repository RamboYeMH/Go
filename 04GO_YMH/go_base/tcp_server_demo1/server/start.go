package server

import (
	"fmt"
	"net"
	"tcp-server-demo1/frame"
	"tcp-server-demo1/metrics"
	"tcp-server-demo1/packet"
)

func handlerPacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode error:", err)
		return
	}
	switch p.(type) {
	case *packet.Submit:
		submit := p.(*packet.Submit)
		fmt.Printf("recv submit: id = %s, payload = %s\n", submit.ID, string(submit.PayLoad))
		submitAck := &packet.SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error:", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}

func handleConn(c net.Conn) {
	metrics.ClientConnected.Inc() // 连接建立 +1
	defer func() {
		metrics.ClientConnected.Desc()
		c.Close()
	}()
	frameCodec := frame.NewMyFrameCodec()
	for {
		// decode the frame to get the payload
		framePayload, err := frameCodec.Decode(c)
		if err != nil {
			fmt.Println("handleConn: frame decode error:", err)
			return
		}
		metrics.ReqRecvTotal.Add(1) // 收到并解码一个消息请求，ReqRecvTotal消息计数器加1
		// 解包
		ackFramePayload, err := handlerPacket(framePayload)
		if err != nil {
			fmt.Println("handlerConn: handle packet error:", err)
			return
		}
		// 回包
		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error:", err)
			return
		}
		metrics.RspSendTotal.Add(1) // 返回响应后，RspSendTotal消息计数器减1
	}
}

func Connection() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		go handleConn(c)
	}
}
