package main

import "fmt"

func main() {
	// 创建map的几种方法
	m := map[string]string{
		"name": "lvtu",
		"sex": "male",
		"age": "27",
		"add": "guangzhou",
	}
	m2 := make(map[string]int)	// m2 == empty map
	var m3 map[string]int	// m3 == nil ,可以参与运算的
	fmt.Println(m, m2, m3)
	// 遍历map
	for k, v := range m{
		fmt.Println(k, v)
	}
	// 获得年龄
	age, ok := m["age"]
	fmt.Println(age, ok)

	if agr, ok := m["agr"]; ok {
		fmt.Println(agr)
	}else {
		fmt.Println("key does not exist")
	}

	// 删除值
	sex, ok := m["sex"]
	fmt.Println(sex, ok)
	delete(m, "sex")
	sex, ok = m["sex"]
	fmt.Println(sex, ok)
}
