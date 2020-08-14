/*
	interface类型默认是一个指针（引用类型），如果没有初始化就使用，那么会输出nil
	一个接口如果继承其他接口，就要实现其他接口的所有方法
*/
package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}

type stu struct {
}
func (s stu) test01() {
}
func (s stu) test02() {
}
func (s stu) test03() {
}

// 空接口,没有任何方法，所以所有类型都实现了空接口
// 可以把任何一个变量赋给空接口
type T interface {
}

func main() {
	var s stu
	var a AInterface = s
	a.test01()
	a.test02()
	a.test03()

	// 空接口接收任何数据类型,s实现了空接口
	var t T = s
	fmt.Println(t)
	// 空接口 另一种写法
	var num float64 = 8.8
	var t2 interface{} = s
	t2 = num
	t = num
	fmt.Println(t2, t)
}
