package main

import (
	"fmt"
	"io/ioutil"
)

/*
	读取文件的另一种方式（一次性把文件全部读到内存），
	适合文件不太大的情况
*/
func main() {
	// 使用ioutil.ReadFile一次性读取到位
	file := "d:/test.txt"
	content, err := ioutil.ReadFile(file) // 没有打开文件，也就不用关闭
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	// 把读取到的内容显示到终端
	fmt.Printf("%v", content) // []byte
	fmt.Printf("%v", string(content)) // 将byte切片转化为字符串
	// 因为没有显示的open文件，因此也不用显示的close文件
	// 因为，文件的open和close被封装到 ReadFile 函数内部
}
