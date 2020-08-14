package main

import (
	real2 "awesomeProject/retriever/real"
	"fmt"
)

// Retriever是一个接口类型，获取url并返回string
type Retriever interface {
	Get(url string) string
}
// 定义一个download方法，传入接口
// 可以让这个接口去完成一个功能
func download(r Retriever) string{
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever  // 创建接口实体
	// r = baidu.Retriever{"this is a fake baidu.com"}

	// 创建自定义结构体类型Retriever实体
	r = real2.Retriever{}

	// 调用download方法传入 自定义结构体实例，这个实例就可以完成接口的方法
	fmt.Println(download(r))
}
