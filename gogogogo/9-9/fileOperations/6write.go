package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 打开一个存在的文件，再原来的内容追加内容 “ABC!ENGLISH”
*/

func main() {
	// 1.打开已经存在的文件 abc.txt
	filePath := "d:/abc.txt"
	// os.O_WRONLY：写的方式打开。os.O_APPEND：写文件时将数据追加到内容的末尾
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	// 及时关闭file句柄
	defer file.Close()
	// 准备写入5句 “hello dog”
	str := "ABC!ENGLISH!\r\n"
	// 写入时，使用带缓存的 *writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer是带缓存，因此在调用writeString方法时，
	// 其实内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	// 真正写入到文件中，否则文件中会没有数据！！！
	writer.Flush()
}
