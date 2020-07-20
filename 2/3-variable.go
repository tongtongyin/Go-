package main

import (
	"fmt"
	"os"
)

func main() {
	var a string = "hello"
	var b = "hello"	// 省略类型，自动推导类型
	var c string	// 省略值或表达式，自动初始化变量值
	// a="hello", b="hello", c=""
	fmt.Printf("a=%q, b=%q, c=%q\n", a, b, c)

	var i, j, k int	// 声明同一类型的多个变量
	var d, f, s = true, 2.3, "love"	// 不加类型，自动由初始化表达式推到
	fmt.Println(i,j,k,d,f,s)

	// 一组变量可以通过调用一个函数，由函数返回的多个返回值初始化
	var g, err = os.Open("./abc.txt")
	fmt.Println(g, err)

	shortVariableState()
}

// 简短变量声明
func shortVariableState(){
	t := 0.0
	i, j := 0, 1	// 简短变量声明
	i, j = j, i	//交换i和j的值
	// 用函数的返回值初始化变量
	f, err := os.Open("./abc.txt")
	// 只声明了一个out变量，err之前声明过，现在只是赋值操作
	// 简短变量声明中，至少要声明一个新的变量，不可以全是之前的变量
	// 如果都是重复的变量，就变成了赋值 =
	out, err := os.Create("./def.txt")
	fmt.Println(t, i, j, f, err, out)
}
