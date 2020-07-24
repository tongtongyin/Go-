package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
// 循环练习
func main() {
	//fmt.Println(sum(100))
	//fmt.Println(converToBinary(13))
	printFile("abc.txt")
}
// 求1到100的和
func sum (n int)  int{
	resutlt := 0
	for i := 1; i <= n; i++ {
		resutlt += i
	}
	return resutlt
}
// 整数转化为二进制
func converToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {	// 可以没有初始条件
		lsb := n % 2 // 每次取最低位
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)	// 停下程序，报错
	}
	scanner := bufio.NewScanner(file)
	// 初始条件和递增条件都可以省略，相当于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}