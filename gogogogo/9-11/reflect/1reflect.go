package main

import (
	"fmt"
	"reflect"
)

/*
 反射的使用场景
 1.结构体标签的应用
 2.使用反射机制，编写函数的适配器（桥连接）
 3....

 基本介绍：
 1.反射可以在运行时动态获取变量的各种信息，比如变量的类型(type)，类别(kind)
 2.如果是结构体变量，还可以获取到结构体本身的信息(包括字段，方法)
 3.通过反射，可以修改变量的值，可以调用关联的方法
 4.使用反射，需要 import "reflect"
*/

// 演示基本数据类型反射
func reflectTest01(b interface{})  {
	// 通过反射获取传入变量的type，kind,值
	// 1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType =",rType)
	// 2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v rVal type=%T \n",rVal, rVal)

	n := 2 + rVal.Int()
	fmt.Println("n=",n)
	// 3.将rVal转成 interface{}
	iv := rVal.Interface()
	// 将interface{}通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2=",num2)
}
// 演示结构体的反射
type Student struct {
	Name string
	Age int
}
func reflectTest02(b interface{})  {
	// 通过反射获取传入变量的type，kind,值
	// 1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType =",rType)
	// 2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v rVal type=%T \n",rVal, rVal)
	// 3.将rVal转成 interface{}
	iv := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iv, iv)
}

func main() {
	// 演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作
	// 1.定义一个int
	// var num int = 100
	// reflectTest01(num)
	// 2.定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age: 20,
	}
	reflectTest02(stu)
}
