package main

import (
	"fmt"
	"runtime"
	"time"
)
/*
 goroutine是一种协程Coroutine：轻量级“线程”
 非抢占式多任务处理，由协程主动交出控制权，什么时候想交出控制权或者不交出控制权，由协程内部主动决定的
 编译器/解释器/虚拟机成眠的多任务，不是操作系统的多任务
 编译器会把go func解释成一个协程，具体执行时go语言有一个调度器去调度协程
 多个协程可能在一个或多个线程上运行，由调度器决定
 线程：抢占式多任务处理，没有控制权，做到一半，哪怕一个语句执行到一半都会被操作系统从中间掐掉，转到其他线程做，然后操作系统还会回来这个线程继续做）
*/

func main() {
	var a [10]int
	for i:=0; i<10; i++ {
		go func(i int) {
			for {
				// fmt.Printf("Hello World form goroutine %d \n", i)
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
