package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
打开一个存在的文件，将原来的内容读出显示在终端，
并且追加5句 “hello,北京！”
*/
func main() {
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	defer file.Close()
	// 先读取原来文件的内容，并显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// 显示到终端
		fmt.Print(str)
	}
	// 写入
	str := "hello,北京！\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
