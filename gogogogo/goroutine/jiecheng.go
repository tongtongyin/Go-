package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	lock sync.Mutex
)

func factorial(n int) {
	temp := 1
	for i:=1; i<=n; i++ {
		temp *= i
	}
	lock.Lock()		// 加锁
	myMap[n] = temp
	lock.Unlock()	// 解锁
}

func main() {
	for i:=1; i<=10; i++ {
		go factorial(i)
	}
	time.Sleep(3 * time.Second)
	// 读map的时候为什么不加锁会出现资源竞争，
	// 因为程序员自己知道3秒可以写完map，但主程序并不知道，
	// 因此底层可能仍出现资源竞争，所以读的时候加互斥锁课解决竞争问题
	lock.Lock()	// 加锁
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()	// 解锁
}
