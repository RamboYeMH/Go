package frame

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"testing"
)

type FramePayload []byte

// StreamFrameCodec 接口类型有两个方法Encode与Decode.
type StreamFrameCodec interface {
	Encode(writer io.Writer, payload FramePayload) error //data -> frame,并写入io.Writer
	Decode(io.Reader) (FramePayload, error)              // 从io.Reader中提取frame payload，并返回给上层
}

var ErrShortWrite = errors.New("short write")
var ErrShortRead = errors.New("short read")

type myFrameCodec struct {
}

func (p *myFrameCodec) Encode(w io.Writer, f FramePayload) error {
	var totalLen int32 = int32(len(f)) + 4
	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}
	n, err := w.Write(f)
	if err != nil {
		return err
	}
	if n != len(f) {
		return ErrShortWrite
	}
	return nil
}

// 先从io.Reader读取四字节的数据，获取到当前frame的长度，再根据长度读取对应的数据
// 这样就可以保证每次读取的包是自己想要的

func (p *myFrameCodec) Decode(r io.Reader) (FramePayload, error) {
	var totalLen int32
	// binary.Read 或 write会根据参数的高度，读取或写入对应的字节个数的字节
	// 这里totalLen使用int32,那么Read或Write只会读操作流中的4个字节
	// 网络字节序使用大端字节序（BigEndian），因此无论是 Encode 还是 Decode，我们都是用 binary.BigEndian；
	err := binary.Read(r, binary.BigEndian, &totalLen)

	if err != nil {
		return nil, err
	}
	buf := make([]byte, totalLen-4)
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	if n != int(totalLen-4) {
		return nil, ErrShortRead
	}
	return buf, nil
}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}

//region ReturnErrorWriter

type ReturnErrorWrite struct {
	W  io.Writer
	Wn int // 第几次调用Write返回错误
	wc int // 写操作次数计数
}

func (w *ReturnErrorWrite) Write(p []byte) (n int, err error) {
	w.wc++
	if w.wc >= w.Wn {
		return 0, errors.New("write error")
	}
	return w.W.Write(p)
}

type ReturnErrorReader struct {
	R  io.Reader
	Rn int // 第几次调用Read返回错误
	rc int // 读操作次数计数
}

func (r *ReturnErrorReader) Read(p []byte) (n int, err error) {
	r.rc++
	if r.rc >= r.Rn {
		return 0, errors.New("read error")
	}
	return r.R.Read(p)
}

func TestEncodeWithWriteFail(t *testing.T) {
	codec := NewMyFrameCodec()
	buf := make([]byte, 0, 128)
	w := bytes.NewBuffer(buf)
	// 模拟binary.Write返回错误码
	err := codec.Encode(&ReturnErrorWrite{
		W:  w,
		Wn: 1,
	}, []byte("hello"))
	if err == nil {
		t.Errorf("want non-nil, actual nil")
	}
	// 模拟w.Write返回错误
	err = codec.Encode(&ReturnErrorWrite{
		W:  w,
		Wn: 2,
	}, []byte("hello"))
	if err == nil {
		t.Errorf("want non-nil, actual nil")
	}
}

func TestDecodeWithReadFail(t *testing.T) {

}

// endregion
