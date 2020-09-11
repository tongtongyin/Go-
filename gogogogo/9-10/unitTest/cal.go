package main

func addUpper(n int) int {
	res := 0
	for i := 1; i <=n; i++ {
		res += i
	}
	return res
}

// 求两个数的差
func getSub(n1, n2 int) int {
	return n1 - n2
}