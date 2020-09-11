/*
  统计1-8000之间，哪些是素数。利用goroutine和channel
*/
package main

import (
	"fmt"
	"time"
)

// 2.将8000个数放入intChan中
func putNum(intC chan int)  {
	for i := 0; i < 150000; i++{
		intC <- i
	}
	close(intC)
}
// 3.从intChan中取数据，并判断是否为素数，如果为素数放入primeChan中,如果放完了，在管道exitChan中写True
func primeNum(intC chan int, primeC chan int, exitC chan bool) {
	for {
		// time.Sleep(time.Millisecond)
		num, ok := <- intC // 判断intC中是否取完
		if !ok {
			break
		}
		// 判断是否为素数
		flag := true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		// 将素数写入primeC
		if flag {
			primeC <- num
		}
	}
	fmt.Println("有一个primeNum协程因取不到数据，退出")
	// 这里不能关闭primeC，因为其他协程可能还未写完数据到primeC
	// 如果此协程写完数据到primeC，则像exitC中写入true
	exitC <- true
}

func main() {
	// 1.首先创建三个channel；intChan,primeChan,exitChan
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()
	// 4.开启一个协程，将8000个数放入到intChan中
	go putNum(intChan)
	// 5.开启4个协程，取出num，判断素数
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	// 6.开启一个协程,用匿名方法判断exitChan中是否已经写完4个True，如果写完，则关闭primeChan
	go func () {
		for i := 0; i < 4; i++ {
			<- exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程时间耗时=", end - start)
		// 当从exitChan取出4个true，即可关闭primeChan
		close(primeChan)
	}()

	// 7.遍历primeChan,查看素数
	for {
		_, ok := <- primeChan
		if ok {
			// fmt.Printf("素数：%d\n", res)
		}else {
			break
		}
	}
	fmt.Println("主线程退出")
}
