package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
 统计一个文件中含有的英文、数字、空格及其他字符数量。
*/
// 定义一个结构体，用于保存统计的结果
type CharCount struct {
	chCount int // 记录英文个数
	NumCount int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	otherCount int // 其他字符的个数
}

func main() {
	// 思路: 打开一个文件，创一个Reader
	// 每读取一行，就去统计改行有多少个英文，数字，空格
	// 将结果保存到一个结构体
	fileName := "d:/count.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	// 定义CharCount 实例
	var count CharCount
	// 创建一个reader
	reader := bufio.NewReader(file)
	// 开始循环读取fileName内容
	for {
		str, err := reader.ReadString('\n')
		// 遍历 str，进行统计
		for _, v := range str {
			switch {
				case v >= 'a' && v <= 'z':
					fallthrough	// 穿透处理
				case v >= 'A' && v <= 'Z':
					count.chCount++
				case v == ' ' || v == '\t':
					count.SpaceCount++
				case v >= '0' && v <= '9':
					count.NumCount++
				default:
					count.otherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}
	// 输出统计结果
	fmt.Printf("字符个数=%v 数字个数=%v 空格个数=%v 其他字符=%v",
		count.chCount, count.NumCount, count.SpaceCount, count.otherCount)
}
