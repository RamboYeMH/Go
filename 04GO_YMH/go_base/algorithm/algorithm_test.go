package algorithm

import (
	"fmt"
	"testing"
)

var source = []int{1, 67, 5, 4, 12, 45, 94, 415, 5}

func TestBubbleSort(t *testing.T) {
	SelectSort(source)
	fmt.Println(source)
}
