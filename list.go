package main

import "errors"

type List struct {
	First *Node
	Last  *Node
	Size  int
}

func NewList() *List {
	return &List{
		First: nil,
		Last:  nil,
		Size:  0,
	}
}

func (l *List) isInitialized() bool {
	return l.First != nil && l.Last != nil
}

func (l *List) initialize(node *Node) {
	l.First = node
	l.Last = l.First
	l.First.Next = l.Last
}

func (l *List) push(value interface{}) {
	node := NewNode(value)

	if l.isInitialized() {
		l.Last.append(node)
		l.Last = node
	} else {
		l.initialize(node)
	}

	l.Size = l.Size + 1
}

func (l *List) poll() (*Node, error) {
	node := l.First

	if l.First == nil {
		return nil, errors.New("list is empty")
	}

	l.First = l.First.Next
	l.Size = l.Size - 1

	if l.Size == 0 {
		l.First = nil
		l.Last = nil
	}

	return node, nil
}

func (l *List) forEach(fn func(interface{})) {
	for node := l.First; node != nil && l.Size != 1; node = node.Next {
		fn(node)
	}
}
