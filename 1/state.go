package main

import "fmt"
// 包一级范围声明一个常量，可在整个包对应的每个源文件中访问
// 不仅仅在其声明语句所在的源文件中访问,例如再e1.2.go中也可以访问boilongF
const boilongF = 212.0

// 函数声明由一个函数名、参数列表（由函数的调用者提供参数变量的具体值）、
// 、一个可选的返回值列表和包含函数定义的函数体组成。
// 如没有返回值，则返回值列表是省略的。
// 执行函数从第一个语句开始，直到遇到return，
// 如没有return，则执行到函数末尾，然后返回到函数调用者。
func main() {
	// main函数内部声明的变量，只能再该函数下被访问
	var f = boilongF
	var c = (f - 32) * 5/9
	// 打印出212°F or 100°C
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
