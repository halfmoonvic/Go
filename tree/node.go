package tree

import "fmt"

// 类似于 ts 中定义一个 接口
// 指明这个对象都有哪些属性
type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// 类似于 js 中 定义一个对象的 有哪些属性，只不过这个属性值是一个方法
func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int) {
	// nil 指针也可以调用方法
	// 即 nil.setValue 这个语句也是行得通的
	if node == nil {
		fmt.Println("Setting value to node, Ignored")
		return
	}
	(*node).Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	// fmt.Println(node)
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
