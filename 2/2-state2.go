package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	// 格式化输出 "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	// "212°F = 100°C"
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(boilingF))
}

// 函数名：fToC
// 参数列表：(f float64)->只有一个参数f，float64为参数的类型
// 返回值类型 float64
// 函数体就一句话 return (f - 32) * 5 / 9
func fToC(f float64) float64{
	return (f - 32) * 5 / 9
}
