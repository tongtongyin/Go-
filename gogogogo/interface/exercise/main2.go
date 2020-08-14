package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 定义英雄类型的结构体
type Hero struct {
	name string
	age int
}
// 定义英雄类型的切片，将每个英雄装入
type heroSlice []Hero

// 用这个英雄类型的切片完成sort结构，
// 使sort接口可以对这个切片排序
func (hs heroSlice) Len() int {
	return len(hs)
}
func (hs heroSlice) Less(i, j int) bool {
	return hs[i].age < hs[j].age
}
func (hs heroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}


func main() {
	// 给切片里添加hero
	var heroes heroSlice
	for i:=0; i<10; i++ {
		h := Hero{
			name: fmt.Sprint("英雄-%d", rand.Intn(100)),
			age: rand.Intn(100),
		}
		heroes = append(heroes, h)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	sort.Sort(heroes)
	fmt.Println("===排序后===")
	for _, v := range heroes {
		fmt.Println(v)
	}

	
}
