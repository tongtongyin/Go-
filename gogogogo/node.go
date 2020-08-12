/*
	go语言仅支持封装，不支持继承和多态
	没有class，只有struct
*/
/*
	要改变内容必须使用指针接收者
	结构过大也考虑使用指针接收者（性能问题）
	一致性：如有指针接收者，最好都是指针接收者
	值接收者是go语言特有
	值/指针 接收者均可接收 值/指针
*/
package main
import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

// 给结构体定义方法
// (node treeNode) 接收者
// 调用该方法时，可以通过treNode实例来.print
func (node treeNode) print() {
	fmt.Print(node.value)
}
func print2(node treeNode)  {
	fmt.Print(node.value)
}

// 该方法改不掉node节点中的值，因为 (node treeNode)是值传递
func (node treeNode) setValue(value int)  {
	node.value = value
}
// 必须传指针才能改掉node节点中的值；(node *treeNode)指针接收者
func (node *treeNode) setValue2(value int)  {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value
}

// 工厂函数，返回一个结构的地址
// 创建值为value的treeNode，返回这个节点的地址
func createNode(value int) *treeNode {
	// treeNode{value: value}是一个局部变量，
	// go语言中，可以返回局部变量的地址给外部用
	return &treeNode{value: value}
}

// 中序遍历;接收者为(node *treeNode)
func (node *treeNode) traverse() {
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	// 创建节点的几种方式
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	// 不论地址还是结构本身，一律使用.来访问成员
	root.right.left = new(treeNode)	// root.right是一个指针，但是也可以直接.left。不管是指针还是实例都可以一路.下去
	root.left.right = createNode(2)

	// 存放节点的切片
	//nodes := []treeNode {
	//	{value: 3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)

	root.print()
	// print2(root) // 用这个方法效果是一样的

	root.right.left.setValue2(4) // 指针接收；将root.right.left这个节点的地址给方法
	root.right.left.print()	// 将指针解析为内容，拷贝到print方法

	//var pRoot *treeNode // 声明treeNode类型的指针变量,nil
	//pRoot.setValue2(200)
	//pRoot = &root		// pRoot保存了root的地址
	//pRoot.setValue2(300)	// pRoot将root中的值改为300
	//pRoot.print()

	fmt.Println()
	root.traverse()
}
