package main

import "fmt"

func main() {
	x := 1
	p := &x	// p初始化为x的指针（地址）； *int类型
	fmt.Println(*p)	// 取地址中的值-> 1
	*p = 2 // 将地址中的值由1改为2
	fmt.Println(x)	// 2
	pointerTest()
	var point = f() // 局部变量的地址
	pointValue := *point // 地址对应的值
	fmt.Println(point, pointValue)
	// false， 因为每调用一次函数都创建一个局部变量，这两个局部变量的地址不同
	fmt.Println(f() == f())

	v := 1
	w := incr(&v) // w保存 2
	// 因为先执行incr(&v),v的值变为3，然后打印 w=2, v=3, incr(&v)=3
	fmt.Println(w, v, incr(&v)) // 2 3 3
	fmt.Println(incr(&v))	// 4
}

// 初始化变量的指针测试
func pointerTest(){
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
}

// 返回局部变量的指针
func f() *int{
	v := 1
	return &v
}

// 参数为一个整形变量的地址，
// 函数功能：将该地址的变量值自增1，返回自增后的值
func incr(p *int) int{
	*p++
	return *p
}