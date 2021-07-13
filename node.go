package main

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

func (n *Node) hasNext() bool {
	return n.Next != nil
}

func (n *Node) append(node *Node) {
	n.Next = node
}
