package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)	//定义字典数据结构存放每个句子和对应出现次数
	files := os.Args[1:]	// 保存输入的字符串到files->['a', 'ab', 'abc',...]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {		// _列表索引下标，arg列表元素
			// 返回f->打开的文件，被Scanner读，err->错误信息，如果err==nil文件打开成功，否则文件出错。
			f, err := os.Open(arg)
			// 如果文件打开时出错打印错误信息，continue掉这个文件
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// 将文件中的行加入字典，并统计每行出现的次数
			countLines(f, counts)
			f.Close()	//关闭文件
		}
	}
	// 打印出有重复的行及 对应的次数
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
// 传入文件，字典，将文件中的每一行加入字典中，key->行，value->该行出现的次数
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] >1{	 // 练习1.4 出现重复的行时打印文件名称
			fmt.Println(f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
