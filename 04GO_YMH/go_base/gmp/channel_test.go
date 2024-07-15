package gmp

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func ChannelTest(T *testing.T) {
	// var ch chan string 默认为nil
	//ch1 := make(chan int)    // 无缓冲区channel
	//ch2 := make(chan int, 5) // 有缓冲区channel
	//ch1 <- 13                // 将整型字面量13发送到无缓存区的channel类型ch1中
	//n := <-ch1               // 从无缓冲区channel类型变量ch1中接受一个整型值存储到整型变量中
	//ch2 <- 17                // 将整型字面量17发送到带缓存区channel类型变量ch2中
	//m := <-ch2               // 从带缓存channel类型变量ch2中接收一个类型值存储到整型变量m中
}

// 只能写的chan
func product(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
}

// 只能读的chan
func consumer(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func start() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		product(ch)
		wg.Done()
	}()

	go func() {
		consumer(ch)
		wg.Done()
	}()
	wg.Wait()
}

/**
无缓冲chan来作用于新号传递
*/

type signal struct {
}

func worker() {
	println("worker is working....")
	time.Sleep(time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work....")
		f()
		c <- signal{}
	}()
	return c
}

func spawnMain() {
	println("start work")
	c := spawn(worker)
	<-c
	println("worker work done!")
}

func workerGroup(i int) {
	fmt.Printf("worker %d: is working....\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d: woks done \n", i)
}

// 无缓冲chan实现1对n的新号通知,返回一个值读chan;
// 如果谁拿到 这个返回值 且c <- signal{}没被执行，那么这个协程将被阻塞
func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d: start to work... \n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}
	go func() {
		wg.Wait()
		c <- signal{}
	}()

	return c
}

func groupStart() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(workerGroup, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	// 无缓冲的多个接收者，如果发消息 只有一个接收者能收到其他的就有死锁
	// 如果用close就给全体发消息
	close(groupSignal)
	<-c
	fmt.Println("the group of workers work done!")
}

func TestGroup(t *testing.T) {
	groupStart()
}

// 基于替代锁机制
// 传统模式基于 “传统模式” + "互斥锁"的Goroutine安全的计数器的实现：
type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func TestTraditional(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine - %d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type newCounter struct {
	c chan int
	i int
}

func NewCounter() *newCounter {
	cter := &newCounter{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()

	return cter
}

func (cter *newCounter) Increase() int {
	return <-cter.c
}
