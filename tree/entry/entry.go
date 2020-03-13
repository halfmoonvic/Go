// 使用组合

package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	*tree.Node
}

func (mtn *myTreeNode) postOrder() {
	if mtn == nil || mtn.Node == nil {
		return
	}
	left := myTreeNode{mtn.Left}
	left.postOrder()
	right := myTreeNode{mtn.Right}
	right.postOrder()
	mtn.Print()
}

func main() {
	// var root tree.Node

	// root = tree.Node{Value: 3}
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	fmt.Println("\n")
	// myRoot := myTreeNode{&root}
	root.postOrder()

	// pRoot := &root
	// pRoot.Print()

}
