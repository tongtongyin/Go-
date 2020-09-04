package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int{
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n",
				id, <-c)
		}
	}()
	return c
}
func chanDemo()  {
	// 建了一个数组，数组每个元素都是一个channel
	var channels [10]chan<- int
	for i:=0; i<10; i++ {
		channels[i] = createWorker(i)
	}
	// 每个channel中都有一个英文字母
	for i:=0; i<10; i++ {
		channels[i] <- 'a' + i
	}
	for i:=0; i<10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
func main() {
	chanDemo()
}
