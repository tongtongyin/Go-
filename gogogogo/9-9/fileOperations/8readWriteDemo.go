package main

import (
	"fmt"
	"io/ioutil"
)

/*
 将d:/abc.txt文件内容导入到 d:/kkk.txt
*/
func main() {
	// 1.将d:/abc.txt文件内容读取到内存
	// 2.将读取到的内容写入 d:/kkk.txt

	file1Path := "d:/abc.txt"
	file2Path := "d:/kkk.txt"
	data, err1 := ioutil.ReadFile(file1Path)
	if err1 != nil {
		fmt.Printf("read file error = %v \n", err1)
		return
	}

	err2 := ioutil.WriteFile(file2Path, data, 0666)
	if err2 != nil {
		fmt.Printf("write file err=%v \n", err2)
	}
}
