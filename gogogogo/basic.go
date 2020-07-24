package main
import "fmt"
import "math"

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}
func variableInitValue(){
	var a, b int = 3, 4
	var s string = "abd"
	fmt.Println(a, b, s)
}
func variableTypeDeduction(){
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b ,c, s)
}
func variableShort(){
	a, b, c, s := 6, 7, true, "def"
	fmt.Println(a, b, c, s)
}
// 强制类型转换
func triangle(){
	var a,b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}
// 常量类型
func consts(){
	const filename = "abc.txt"
	const a, b = 3, 4 // 不指定类型，下面的Sqrt函数就会当成float来用
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
// 枚举类型
func enums(){
	const(
		java = iota		//自增
		_
		cpp
		golang
		javascript
	)
	const(
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb

	)
	fmt.Println(java, cpp, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(aa, ss, bb)
	triangle()
	consts()
	enums()
}