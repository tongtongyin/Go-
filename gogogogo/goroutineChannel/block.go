package main

import (
	"fmt"
)

// 3.写入channel
func writeData2(intC chan int)  {
	for i := 1; i <= 50; i++ {
		intC <- i
		fmt.Println("write", i)
	}
	close(intC)
}
// 4.读取channel
func readData2(intC chan int, exitC chan bool)  {
	for {
		v, ok := <- intC
		if !ok {	// 如果管道里没有数据了，退出循环
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	// 已经读完数据，将exitC通道中写入true
	exitC <- true
	close(exitC)	// 关闭通道
}
func main() {
	// 1.建立读写channel
	intChan := make(chan int, 50)
	// 2.建立退出channel
	exitChan := make(chan bool, 1)
	// 5.并行执行读写
	go writeData2(intChan)
	// 只写不读会造成阻塞。但是读写的频率不一致，无所谓
	// go readData2(intChan, exitChan)
	// 6.读完的时候会塞进管道exitChan一个布尔值
	// 循环的读取这个bool值，直到取出来之后，exitChan管道为空，退出循环
	for {
		_, ok := <- exitChan // 循环执行的将exitChan中的布尔值取出来
		if !ok {	// 直到exitChan中为空，退出循环
			break
		}
	}
}
