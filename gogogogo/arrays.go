package main

import "fmt"

func printArrays(arr [5]int)  {
	arr[0] = 100
	for i, v := range arr{
		fmt.Println(i, v)
	}
}
// 传递数组指针，就可以修改原数组里面的值
func printArrays2(arr *[5]int)  {
	arr[0] = 100
	for i, v := range arr{
		fmt.Println(i, v)
	}
}
// go语言一般不使用数组
func main() {
	// 定义数组的几种方法
	// 数量写在类型前
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2,4,6,8,10}

	var grid [4][5]int
	fmt.Println(arr1,arr2,arr3,grid)

	// 两种遍历方式
	for i:=0; i<len(arr3); i++{
		fmt.Println(arr3[i])
	}
	// range遍历得到数组下标，只要一个下标
	for i:= range arr3{
		fmt.Println(arr3[i])
	}
	// range 遍历数组下标和值
	for i, v := range arr3{
		fmt.Println(i, v)
	}
	// 可以通过下划线省略变量
	for _, v := range arr3{
		fmt.Println(v)
	}
	// 可以记录数组的最大值，和对应编号
	maxi := -1
	maxValue := -1
	for i, v := range arr3{
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println(maxi, maxValue)
	// 对数组求和
	sum := 0
	for _, v := range arr3{
		sum += v
	}
	fmt.Println(sum)
	fmt.Println("printArrays(arr1)")
	// 调用这个函数的时候，会对arr1数组中的元素进行一份拷贝，arr1数组本身的值没有变
	printArrays(arr1)
	// printArrays(arr2)  数组长度不对
	fmt.Println("printArrays(arr3)")
	printArrays(arr3)
	fmt.Println("arr1 and arr3")
	// 值类型，在函数里可以改，但是在这里数组的值不会变
	fmt.Println(arr1, arr3)

	fmt.Println("printArrays2(arr1)")
	// 传递数组的指针
	printArrays2(&arr1)
	fmt.Println("printArrays2(arr3)")
	printArrays2(&arr3)
	fmt.Println("arr1 and arr3")
	// 由于传的是数组的指针所以就会改变原数组中的值
	fmt.Println(arr1, arr3)
}
