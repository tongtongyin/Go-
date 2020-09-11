package main

import "fmt"

/*
	案例：
*/

// 在函数内部，ch管道是只写的，可以防止误操作
func send(ch chan <- int, exitChan chan struct{})  {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}
// 在recv函数内部，ch是只读的，防止误操作
func recv(ch <- chan int, exitChan chan struct{})  {
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}
func main() {
	var ch chan int	// ch双向管道
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)

	go send(ch, exitChan) // send协程中的ch是只写的，可以防止误操作
	go recv(ch, exitChan) // recv协程中的ch是只读的，防止误操作

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束...")
}
