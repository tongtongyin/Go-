package main

import (
	"fmt"
	"sort"
)



func main() {
	// 定义数组切片
	var intSlice = []int{0, -1, 10, 7, 90}
	// 使用系统方法对切片进行排序
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}
