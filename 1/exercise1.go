package main

import (
	"fmt"
	"os"
)
// 修改echo程序，使其打印每个参数的索引和值，每个一行。
func main() {
	var s, sep string
	// 修改索引从0开始，即可打印被执行命令本身的名字
	for i := 0; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

}
