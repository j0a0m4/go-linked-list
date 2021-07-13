package main

import (
	"testing"
)

func Test_newNodeValue(t *testing.T) {
	node := NewNode("Placeholder Value")

	if got := node.Value; got != "Placeholder Value" {
		t.Errorf("newNode(\"Placehodler Value\").Value = %s; want \"Placeholder Value\"", got)
	}
}

func Test_newNodeNext(t *testing.T) {
	node := NewNode("Placeholder Value")

	if got := node.Next; got != nil {
		t.Errorf("newNode(\"Placehodler Value\").Next = %s; want nil", got)
	}
}

func Test_hasNextFalse(t *testing.T) {
	node := NewNode("Placeholder Value")

	if got := node.hasNext(); got != false {
		t.Errorf("newNode(\"Placehodler Value\").hasNext() = %v; want false", got)
	}

}

func Test_hasNextTrue(t *testing.T) {
	node := NewNode("Placeholder Value")
	node.Next = NewNode("Placeholder Next")

	if got := node.hasNext(); got != true {
		t.Errorf("newNode(\"Placehodler Value\").hasNext() = %v; want true", got)
	}
}

func Test_append(t *testing.T) {
	node := NewNode("Placeholder Value")
	nextNode := NewNode("Placeholder Next")
	node.append(nextNode)

	if got := node.Next; got != nextNode {
		t.Errorf("newNode(\"Placehodler Value\").append(nextNode) = %v; want %v", got, nextNode)
	}
}

func Test_newList(t *testing.T) {
	list := NewList()

	if got := list.First; got != nil {
		t.Errorf("newList().First = %v; want nil", got)
	}

	if got := list.Last; got != nil {
		t.Errorf("newList().Last = %v; want nil", got)
	}

	if got := list.Size; got != 0 {
		t.Errorf("nesList().Size = %v; want 0", got)
	}
}

func Test_newListisInitialized(t *testing.T) {
	list := NewList()
	list.First = &Node{Value: "", Next: nil}
	list.Last = &Node{Value: "", Next: nil}

	if got := list.isInitialized(); got != true {
		t.Errorf("newList().isInitialized() = %v; want true", got)
	}
}

func Test_newListisNotInitialized(t *testing.T) {
	list := NewList()

	if got := list.isInitialized(); got != false {
		t.Errorf("newList().isInitialized() = %v; want false", got)
	}
}

func Test_newListisInitializedFirst(t *testing.T) {
	list := NewList()
	list.First = &Node{Value: "", Next: nil}

	if got := list.isInitialized(); got != false {
		t.Errorf("newList().First = %v; want false", got)
	}
}

func Test_newListisInitializedLast(t *testing.T) {
	list := NewList()
	list.Last = &Node{Value: "", Next: nil}

	if got := list.isInitialized(); got != false {
		t.Errorf("newList().Last = %v; want false", got)
	}
}

func Test_initilizeList(t *testing.T) {
	list := NewList()
	node := &Node{Value: "", Next: nil}

	list.initialize(node)

	if got := list.First; got != node {
		t.Errorf("newList().First = %v; want %v", got, node)
	}

	if got := list.Last.Value; got != node.Value {
		t.Errorf("list.Last.Value = %v; want %v", got, node.Value)
	}

	if got := list.Last.Next; got != list.First.Next {
		t.Errorf("list.Last.Next = %v; want nil", got)
	}
}
func Test_pushWhenIsInitialized(t *testing.T) {
	node := &Node{Value: "", Next: nil}

	list := &List{
		First: &Node{Value: "", Next: node},
		Last:  node,
		Size:  1,
	}

	list.push("New Value")

	if got := list.Last.Next; got != nil {
		t.Errorf("list.Last.Next = %v; want nil", got)
	}

	if got := list.Last.Value; got != "New Value" {
		t.Errorf("list.Last.Value = %v; want \"New Value\"", got)
	}

	if got := list.Size; got != 2 {
		t.Errorf("list.Size = %v; want 2", got)
	}

}

func Test_pushWhenItsNotInitialized(t *testing.T) {
	list := &List{
		First: nil,
		Last:  nil,
		Size:  0,
	}

	list.push("New Value")

	if got := list.First.Value; got != "New Value" {
		t.Errorf("list.First.Value = %v; want \"New Value\"", got)
	}

	if got := list.Last.Value; got != "New Value" {
		t.Errorf("list.Last.Value = %v; want \"New Value\"", got)
	}

	if got := list.Size; got != 1 {
		t.Errorf("list.Size = %v; want 1", got)
	}
}

func Test_pollFromList(t *testing.T) {
	lastNode := &Node{Value: "last", Next: nil}
	firstNode := &Node{Value: "first", Next: lastNode}

	list := &List{
		First: firstNode,
		Last:  lastNode,
		Size:  1,
	}

	got, err := list.poll()

	if got != firstNode {
		t.Errorf("list.poll() = %v; want %v", got, firstNode)
	}

	if got = list.First; got != nil {
		t.Errorf("list.First = %v; want nil", got)
	}

	if got = list.Last; got != nil {
		t.Errorf("list.Last = %v; want nil", got)
	}

	if err != nil {
		t.Errorf("list.poll() = %v; want nil", got)
	}

	if list.Size != 0 {
		t.Errorf("list.Size = %v; want 0", list.Size)
	}
}

func Test_pollFromEmptyList(t *testing.T) {
	list := &List{
		First: nil,
		Last:  nil,
		Size:  0,
	}

	got, err := list.poll()

	if got != nil {
		t.Errorf("list.poll() = %v; want nil", got)
	}

	if err == nil {
		t.Errorf("list.poll() = %v; want != nil", err)
	}

}
