package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func main() {
	var s []int	// 创建slice第一种方法，定义一个空的slice类型s
	fmt.Printf("s=%d\n", s)
	// 打印前100个奇数,可以发现
	for i:=0; i < 100; i++{
		printSlice(s)	// 随着len的增加，直到每次cap装不下了，那么cap就会乘以2进行扩充。
		s = append(s, i * 2 + 1)
	}
	fmt.Println(s)

	s1 := []int{2,4,6,8}	// 创建一个数组，再生成slice
	s2 := make([]int, 16)	//make方法创建长度16的slice
	s3 := make([]int, 10, 32) // 创建长度10，cap为32的slice
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
	fmt.Println("Copying slice")
	copy(s2, s1)	// 将s1拷贝到s2中
	printSlice(s2)

	fmt.Println("Deleteing elements from slice")
	// 删除s2中下标为3的元素,采用append的方式
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	// 删除头元素
	front := s2[0]
	s2 = s2[1:]
	// 删除末尾元素
	tail := s2[len(s2) -1]
	s2 = s2[:len(s2)-1]
	fmt.Printf("front=%d, tail=%d\n",
		front, tail)
	printSlice(s2)
}
