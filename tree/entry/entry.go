package main

import (
	"learngo/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	// fmt.Println(root.Right.Left)

	root.Right.Left.SetValue(4)
	// root.Right.Left.print()
	// root.setValue(100)
	// root.print()
	root.Traverse()

	// var pRoot *Node
	// fmt.Println(pRoot)
	// pRoot.setValue(200)

	// pRoot = &root
	// fmt.Println(pRoot)
	// pRoot.setValue(300)
	// pRoot.print()
}
