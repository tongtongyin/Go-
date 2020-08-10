package main

import "fmt"

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr=", arr)
	s1 := arr[2:6] // 2,3,4,5
	// slice可以向后扩展，但不可以向前扩展
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	// s[i]不可以超过len(s),向后扩展可以超越len(s),但不可以超越底层数组cap(s).
	s2 := s1[3:5] // 5,6
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))
	fmt.Println(s1, s2)
}
