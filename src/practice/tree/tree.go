package main

import "fmt"

type Node struct {
	le *Node
	data interface{}
	ri *Node
}

func NewNode(left,right *Node)  *Node {
	return &Node{left,nil,right}
}

func (n *Node)SetDate(data interface{})  {
	n.data = data
}

func main() {
	root := NewNode(nil,nil)
	root.SetDate("root node")
	a := NewNode(nil,nil)
	a.SetDate("left node")
	b := NewNode(nil,nil)
	b.SetDate("right node")
	root.le = a
	root.ri = b
	fmt.Println(root)

}
