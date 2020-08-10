package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]	// [2 3 4 5]
	s2 := s1[3:5]	// [5 6]
	s3 := append(s2, 10)	// [5 6 10]

	// s4和s5不在view arr，而是一个新的arr
	// 添加元素时如果超越cap，系统会重新分配更大的底层数组
	s4 := append(s3, 11) // [5 6 10 11]
	s5 := append(s4, 12) // [5 6 10 11 12]
	fmt.Printf("s1=%v,s2=%v,s3=%v,s4=%v,s5=%v\n",
		s1, s2, s3, s4, s5)
	// 重新分配更大的底层数组
	fmt.Printf("cap(s2)=%d,cap(s3)=%d,cap(s4)=%d,cap(s5)=%d\n",
		cap(s2),cap(s3),cap(s4), cap(s5))
}
