package main

import (
	"fmt"
	"math"
)
/*
	1.返回值类型写在最后面
	2.可返回多个值
	3.函数作为参数
 */
func main() {
	if result, err := evel(3,4,"*"); err !=nil {
		fmt.Println("Error:", err)
	}else {
		fmt.Println(result)
	}
	q, r := div(13,3)
	fmt.Println(q, r)
	// 调用 apply
	fmt.Println(apply(
		func(a int, b int) int { // 匿名函数
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3,4))

	fmt.Println(sumArgs(1,2,3,4,5))
}
// 加减乘除计算器, 函数返回两个值
func evel(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// 如果有多个返回值，不用的可以 _ 接收
		q, _ := div(a, b)
		return q, nil
	default://错误处理，返回0，并打印错误信息
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}
// 带余除法，函数可以返回多个参数
func div(a, b int) (q, r int) {
	return a / b, a % b
}
// 函数式编程
// op func(int, int) 方法返回 int，传进的参数为a,b
// apply 方法返回 op 函数 为int类型
func apply(op func(int, int) int, a, b int) int {
	// p := reflect.ValueOf(op).Pointer()
	// opName := runtime.FuncForPC(p).Name()
	// fmt.Printf("Calling function %s with args " +
		// "(%d, %d)\n", opName, a, b)
	return op(a, b)
}
// 函数可以传可变参数列表
// 求可变参数列表的和
func sumArgs(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
