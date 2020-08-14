/*
	1.一个自定义类型可以实现两个接口的方法
	2.go接口里面不能有任何变量
*/
package main

import "fmt"

type Ainterface interface {
	say()
}
type Binterface interface {
	hello()
}

type monster struct {
}

func (m monster) say() {
	fmt.Println("say()方法")
}
func (m monster) hello() {
	fmt.Println("hello()方法")
}

func main() {
	var m monster
	var a Ainterface = m
	var b Binterface = m
	a.say()
	b.hello()
}
