package gmp

import "testing"

type event interface {
	start()
	end()
}

type numInfo struct {
	age  int
	name string
}

func (n *numInfo) print() {

}

type taskEvent struct {
	event
	numInfo
}

func TestEvent(t *testing.T) {
	task := taskEvent{}
	task.end()

}
