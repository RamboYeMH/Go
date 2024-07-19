package server

import (
	"fmt"
	"github.com/lucasepe/codename"
	"net"
	"sync"
	"tcp-server-demo1/frame"
	"tcp-server-demo1/packet"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	Connection()
}

func TestClient(t *testing.T) {
	var wg sync.WaitGroup
	var num int = 1
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			startClient(i)
		}(i + 1)
	}
	wg.Wait()
}

// 启动客户端

func startClient(i int) {
	quit := make(chan struct{})
	done := make(chan struct{})
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()
	fmt.Printf("[client %d]: dial ok\n", i)
	// 生成payload(装载器)
	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}
	frameCodec := frame.NewMyFrameCodec()
	var counter int
	go func() {
		// handler ack
		for {
			select {
			case <-quit:
				done <- struct{}{}
			default:
			}
			fmt.Println("hello")
			err := conn.SetReadDeadline(time.Now().Add(time.Minute * 5))
			if err != nil {
				return
			}
			ackFramePayLoad, err := frameCodec.Decode(conn)
			if err != nil {
				if e, ok := err.(net.Error); ok {
					if e.Timeout() {
						continue
					}
				}
				panic(err)
			}
			p, err := packet.Decode(ackFramePayLoad)
			submitAck, ok := p.(*packet.SubmitAck)
			if !ok {
				panic("not submitAck")
			}
			fmt.Printf("[client %d]: the result of submit ack[%s] is %d\n", i, submitAck.ID, submitAck.Result)
		}

	}()

	for {
		// send submit
		counter++
		id := fmt.Sprintf("%08d", counter) //8 byte string
		payload := codename.Generate(rng, 4)
		s := &packet.Submit{
			ID:      id,
			PayLoad: []byte(payload),
		}
		framePayload, err := packet.Encode(s)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[client %d]: send submit id = %s, payload=%s, frame length = %d\n", i, s.ID, s.PayLoad, len(framePayload)+4)
		err = frameCodec.Encode(conn, framePayload)
		if err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
		if counter >= 10 {
			quit <- struct{}{}
			<-done
			fmt.Printf("[client %d]: exit ok", i)
			return
		}
	}
}
