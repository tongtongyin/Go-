package main

import "fmt"

type Usb interface {
	start()
	stop()
}

// 不用非得是结构体类型，其他自定义类型也可以实现接口
type integer int
func (i integer) start() {
	fmt.Println("start i =",i)
}
func (i integer) stop() {
	fmt.Println("stop i =", i)
}

// 实现接口的方法必须实现接口里面所有的方法，否则不行
type Usb2 interface {
	start()
	stop()
	test()
}
// phone实现了接口，就是实现了Usb接口的所有方法
type Phone struct {
}
func (p Phone) start() {
	fmt.Println("手机开始工作了")
}
func (p Phone) stop() {
	fmt.Println("手机停止工作了")
}
// camera实现了接口，实现了Usb接口的所有方法
type Camera struct {
}
func (c Camera) start() {
	fmt.Println("相机开始工作")
}
func (c Camera) stop() {
	fmt.Println("相机停止工作了")
}

type Computer struct {
}
// 编写一个Working方法，接收一个Usb接口类型变量
// 只要实现了Usb接口。就是实现了Usb接口声明的所有方法。
func (c Computer) Working(usb Usb) {
	// 通过usb接口变量来调用start和stop方法
	usb.start()
	usb.stop()
}

func main() {
	// 创建结构体变量
	iphone := Phone{}
	camera := Camera{}
	computer := Computer{}
	// 把iphone传进去，
	// 在Working方法中调用的就是phone的start()方法和stop()方法
	computer.Working(iphone)
	computer.Working(camera)

	// 接口本省不能创建实例
	// 但是实现了该接口的自定义类型实例 可以赋给接口
	var i integer = 10
	var u Usb = i
	u.start()

}
