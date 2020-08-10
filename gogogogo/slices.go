package main

import "fmt"
// 可以传入数组切片，进而修改数组的值
func updateSlice(s []int)  {
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	// 前闭后开
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])

	s1 := arr[2:]
	fmt.Println("s1 =", s1) // 2,3,4,5,6,7
	s2 := arr[:]
	fmt.Println("s2 =", s2)	// 0,1,2,3,4,5,6,7
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)	// 将切片s1传入，修改切片中的值
	fmt.Println(s1)	// 100,3,4,5,6,7
	fmt.Println(arr) // 0,1,100,3,4,5,6,7

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)	// 100,1,100,3,4,5,6,7
	fmt.Println(arr) // 100,1,100,3,4,5,6,7

	fmt.Println("reslice")  // 可以对切片再进行切片
	fmt.Println(s2) // 100,1,100,3,4,5,6,7
	s2 = s2[:5]		// 100,1,100,3,4
	fmt.Println(s2)
	s2 = s2[2:]		// 100,3,4
	fmt.Println(s2)

}
