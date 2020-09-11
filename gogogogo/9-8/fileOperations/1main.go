package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件,file=文件对象=文件指针=文件句柄
	file, err := os.Open("d:/test.txt")
	if err !=nil {
		fmt.Println("open file err=", err)
	}
	// 输出文件，看文件是什么 其实是一个指针
	fmt.Printf("file=%v", file) // file=&{0xc00007e780}
	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
