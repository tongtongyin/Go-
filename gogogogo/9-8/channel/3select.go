package main

import (
	"fmt"
)

/*
	使用select可以解决从管道取数据的阻塞问题
*/

func main() {
	// 1.定义一个管道 10个int数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 2.定义一个管道 5个string数据
	var strChan chan string
	strChan = make(chan string, 10)
	for i := 0; i < 10; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	// 有时候不好确定该管道是否应该关闭了
	// 可以使用select方式可以解决
	label:
	for {
		select {
	//如果intChan一直没有关闭，不会阻塞发生deadlock，会自动到下一个case匹配
			case v := <- intChan:
				fmt.Printf("从intChan读取数据%d\n", v)
			case v := <- strChan:
				fmt.Printf("从intChan读取数据%s\n", v)
			default:
				fmt.Println("取不到了。。。")
				break label
		}
	}
}
