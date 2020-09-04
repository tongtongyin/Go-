package main

import "fmt"

// channel的遍历和关闭
// 关闭：使用内置方法close，关闭后，不能再向channel写数据，但可以读
// 遍历：channel支持for-range遍历，用普通for循环长度会变化
// 	遍历时如果channel没关闭，会出现deadlock错误
// 	遍历时如果channel已关闭，就会正常遍历
func main() {
	// 创建3大小的管道,写入两个数据
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	// 关闭管道，不能写数据，可以读数据
	close(intChan)
	// intChan <- 300
	n := <- intChan
	fmt.Println(n)
	// 遍历管道，不能用普通for
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	// 首先要关闭管道,不关闭也可以遍历出所有数据，但是会出现deadlock。
	// 因为如果不关闭的话，管道最后没有关闭的标志，
	// 遍历的时候以为还可以取管道里的值，就一直在等带，造成死锁
	close(intChan2)
	for v := range intChan2 {
		fmt.Println(v)
	}
}
