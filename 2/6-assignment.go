package main

import "fmt"

func main() {
	// 使用元组赋值
	// i, j, k = 2, 3, 5

	// 用元组来接函数的返回值，函数返回值的数量与元组数量需一致
	// f, err = os.Open('abc.txt')

	// 下划线 _ 来丢弃不需要的值
	// _, err = io.Copy(dst, src)

	// 隐式赋值
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals[0], medals[1], medals[2])
}

// 辗转相除法求最大公约数
func gcd(x, y int) int{
	for y != 0{
		x, y = y, x%y
	}
	return x
}

// 斐波那契额数列
func fib(n int) int{
	x, y := 0, 1
	for i:= 0; i < n; i++{
		x, y = y, x+y
	}
	return x
}
