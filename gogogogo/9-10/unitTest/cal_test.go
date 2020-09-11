package main

import (
	"fmt"
	"testing"

)

/*
 testing框架
 将xxx_test.go的所有文件引入，
 扫描文件里面哪些是以Test开头的方法，然后再调用里面要测试的函数1，2，3...
 main(){
	// 2.调用TestXXX()函数
	}
*/

// 编写要给测试用例，去测试addUpper是否正确
func TestAddUpper(t *testing.T)  {
	// 调用
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v", 55, res)
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v", 55, res)
	}
	// 如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...")
}

func TestHello(t *testing.T)  {
	fmt.Println("TestHello被调用。。。")
}

func TestGetsub(t *testing.T)  {
	res := getSub(10, 7)
	if res != 3 {
		// fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v", 55, res)
		t.Fatalf("getSub(10, 7) 执行错误，期望值=%v 实际值=%v", 3, res)
	}
	// 如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...")
}

/*
 1.测试用例文件名必须以_test.go结尾
 2.测试用例函数必须以Test开头，一般来说就是Test+被测试的函数名
 3.TestAddUpper(t *testing.T)的形参类型必须是*testing.T
 4.一个测试文件中，可以有多个测试函数
 5.运行测试用例指令
	(1) go test
	(2) go test -v
 6.当出现错误时，可以使用t.Fatalf来格式化输出错误信息，并退出程序
 7.t.Logf方法可以输出相应的日志
 8.测试用例函数，并没有放在main函数中，也可以执行，这时测试用例的方便之处
 9.PASS表示测试用例运行成功，FALL表示测试用例运行失败
 10.测试单个文件，一定要带上被测试的源文件
	go test -v cal_test.go cal.go
 11.测试单个方法
	go test -v -test.run TestAddUpper
*/

