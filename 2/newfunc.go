package main

import "fmt"

// 另一个创建变量的方法，调用内建new函数
func main() {
	p := new(int)	// p保存了一个 匿名int类型 的地址
	fmt.Println(*p)	// 打印*P为int类型的初始化值  “0”
	*p = 2			// 设置匿名int变量值为 2
	fmt.Println(*p)	// ‘2’

	// 每次new时候返回的都是一个新的地址
	w := new(int)
	q := new(int)
	fmt.Println(w==q)  // false
}

// 返回一个匿名的int类型地址
func newInt() *int{
	return new(int)
}
// 返回一个int类型地址
func newInt2() *int{
	var d int
	return &d
}
