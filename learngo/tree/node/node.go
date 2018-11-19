package node

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func CreateNode(num int)*Node{
	return  &Node{Value:num}
}

func (v *Node)SetValue(num int){
	v.Value = num
}

func (v Node)Print(){
	fmt.Println(v.Value)
}

func (f *Node)Traverse(){
	if f == nil {
		return
	}
	f.Left.Traverse()
	f.Print()
	f.Right.Traverse()

}