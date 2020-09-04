package main

import "fmt"
/*
	存放任意数据类型的管道
*/
type Cat struct {
	Name string
	Age int
}

func main() {
	// 定义一个存放任意数据类型的管道 3个数据
	// var allChan chan interface{}
	allChan := make(chan interface{}, 3)
	// 像管道中添加数据
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat
	// 从管道中推出两个数据
	<- allChan
	<- allChan
	// 从管道中取出最后一个数据
	newCat := <- allChan
	// newCat type is main.Cat, value is {小花猫 4}
	fmt.Printf("newCat type is %T, value is %v\n",
		newCat, newCat)
	// 报错：在编译层面，newCat是一个interface{}，接口里面是没有字段的
	// fmt.Printf("newCat.Name=%v", newCat.Name)

	// 使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}
