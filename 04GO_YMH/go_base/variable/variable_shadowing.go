package main

import (
	"fmt"
	"time"
)

// 变量遮蔽（variable shadowing）
// 遮蔽代码块 作用域
/*
 // 代码块的嵌套
fun foo() { // 代码块1
	{ // 代码块2
		{ //代码块3

		}
	}
}

// go语言没有枚举，因为设计者认为常量和枚举是有很多相通的地方
*/
const (
	ONE = 1 << iota
	TWO
	THREE
	FOUR = iota
)

var a = 11
var anyType [5]any

func foo(n int) {
	// 变量遮蔽的根本问题就是在内层代码中声明了
	a := 1
	a += n
}

func main() {
	param()
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 由于i v都是
	for i, v := range array {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Printf("key = %d, value = %d\n", i, v)
		}()
	}

	// 通过闭包来打印
	for i, v := range array {
		go func(i, v int) {
			time.Sleep(time.Second * 3)
			fmt.Printf("key = %d, value = %d\n", i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 10)

}

func param() (number int, name string) {
	fmt.Printf("age = %d, name = %s\n", number, name)
	return 0, "nil"
}

func foo1() {
	println("call foo")
	bar()
	println("exit foo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the panic:", e)
		}
		panic("panic occurs in bar")
	}()

	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")
}
func zoo() {
	println("call zoo")
	println("exit zoo")
}

func start() {
	println("start")
	foo1()
	println("end ")

}
