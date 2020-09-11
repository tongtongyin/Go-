package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
/*
 读取文件内容的一种方式（缓冲读取）
*/
func main() {
	file, err := os.Open("d:/test.txt")
	if err !=nil {
		fmt.Println("open file err=", err)
	}
	// 当函数退出时，要及时关闭file,否则会有内存泄漏
	defer file.Close()
	// 创建一个 *reader，是带缓冲的, 默认缓冲区4096字节
	// 可以边读边处理
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行符就结束（一行一行读）
		fmt.Print(str) // 输出内容,先输出在判断是否到文件末尾，否则会少输出一行
		if err == io.EOF { // io.EOF表示文件的末尾 end of file
			break
		}
	}
	fmt.Println("文件读取结束")
}
