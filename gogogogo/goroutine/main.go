package main

import (
	"fmt"
	"runtime"
	"time"
)
// 在主线程

func test()  {
	for i:=1; i<=15; i++ {
		fmt.Printf("test() hello,world %d\n", i)
		time.Sleep(time.Second)
	}
}
// 主线程是一个物理线程，直接作用在cpu上的。是重量级的，非常耗费cpu资源。
// 协程从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对小。
func main() {
	go test() // 开启一个协程
	for i:=1; i<=10; i++ {
		fmt.Printf("main() hello,golang %d \n", i)
		time.Sleep(time.Second)
	}
	fmt.Println(runtime.NumCPU())
}
