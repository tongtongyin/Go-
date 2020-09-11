package main

import "fmt"

/*
	管道可以定义为 只读 或者 只写
	默认情况管道是可读可写的
 */
func main() {
	// 可读可写
	// var chan1 chan int

	// 1.声明只写
	var chan2 chan <- int
	chan2 = make(chan int, 3)
	chan2 <- 20	// 可以写
	// num := <- chan2 // 不可以读
	fmt.Println("chan2=", chan2)

	// 2.声明为只读
	var chan3 <-chan int
	num2 := <-chan3 	// 可以读
	// chan3 <- 30 	// 不可以写
	fmt.Println("num2", num2)
}
