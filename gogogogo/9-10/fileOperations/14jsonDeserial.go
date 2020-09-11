package main

import (
	"encoding/json"
	"fmt"
)
/*
	将结构体、map、切片反序列化
	再反序列化一个json字符串时，要确保反序列化后的数据类型和原来的序列化前的数据类型一致
	如果json字符串时通过程序获取到的，则不需要对 " 转义处理。
*/
type Monster1 struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}
// 反序列化：将json反序列化为结构体
func unmarshalStruct()  {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000.1,\"Skill\":\"牛魔拳\"}"
	var monster1 Monster1
	err := json.Unmarshal([]byte(str), &monster1)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Println("反序列化后", monster1)
}
// 反序列化：将json反序列化map
func unmarshalMap()  {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	// 定义一个map
	var a map[string]interface{}
	// 反序列化
	// 反序列化map，不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后a=%v\n", a)
}
// 反序列化：将json反序列化为切片
func unmarshalSlice()  {
	str := "[{\"address\":\"beijing\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"Mexico\",\"Hawaii\"],\"age\":\"20\",\"name\":\"tom\"}]"
	// 定义一个slice
	var slice []map[string]interface{}
	// 反序列化
	// 反序列化map，不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后a=%v\n", slice)
}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
