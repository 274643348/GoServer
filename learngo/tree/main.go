package main

import "../tree/node"
func main() {
	var root node.Node
	root = node.Node{Value:3}
	root.Left = &node.Node{}
	root.Right = &node.Node{5,nil,nil}
	root.Right.Left = new(node.Node)
	root.Left.Right =  node.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
}
