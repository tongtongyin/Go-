package main

import "fmt"

// channel
// 本质是一个队列，先进先出（FIFO）
// 线程安全，多goroutine访问时，不需要加锁，channel本身就是线程安全的
// 有类型的，一个string的channel只能存放string类型数据

// var 变量名 chan 数据类型
// 举例：
// var intChan chan int   //intChan用于存放int数据
// var mapChan chan map[int]string  // mapChan用于存放map[int]string类型
// var perChan chan Person  // 结构体也可以
// var perChan2 chan * Person  //指针也可以
// 说明：
// channel是引用类型。
// channel必须初始化才能写入数据，即make后才能使用。
func main() {
	// 1.创建一个可以存放3个int类型的管道intChan
	var intChan chan int
	intChan = make(chan int, 3)
	// 2.看看intChan是什么
	fmt.Printf("intChan value is %v address is %p \n", &intChan, intChan)
	// 3.向管道写入数据，不能超过容量
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	<- intChan //
	// 4.看看管道长度和cap(容量)
	fmt.Printf("intChan len=%v and cap=%v\n",
		len(intChan), cap(intChan))
	// 5.从管道中读取数据
	var num2 int
	num2 = <- intChan
	var num3 int
	num3 = <- intChan
	num4 := <- intChan
	fmt.Println(num4)
	fmt.Println(num2, num3, num4)
	fmt.Printf("intChan len=%v and cap=%v \n",
		len(intChan), cap(intChan))
	// 6.在没有使用协程的情况下，如果channel中数据全部取出，再取就会报告 deadlock
	// num5 := <- intChan
	// fmt.Println(num5)
}
