package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
 拷贝文件：将一张图片/电影/mp3拷贝到另外一个文件
 io包-> func Copy()
*/

// 编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	// 通过srcFile，获取reader
	reader := bufio.NewReader(srcFile)

	// 打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer dstFile.Close()

	// 通过dstFile，获取writer
	writer := bufio.NewWriter(dstFile)
	// 调用Flush方法，将缓冲的数据写入到文件中
	writer.Flush()

	return io.Copy(writer, reader)
}

func main() {
	// 将d:/ME.png文件拷贝到d:/I.png
	// 调用CopyFile 完成文件拷贝
	srcFile := "d:/ME.png"
	dstFile := "d:/I.png"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝错误 err=%v", err)
	}

}
