package linkedlist

import (
	"fmt"
)

type List interface {
	IsEmpty() bool
	Size() uint
	Get(uint) *Node
	Insert(position uint, value int)
	Remove(position uint) *Node
}

type Node struct {
	Next  *Node
	Value int
}

type LinkedList struct {
	Head *Node
}

func (l LinkedList) IsEmpty() bool {
	return l.Head == nil
}

func (l LinkedList) Size() uint {
	var size uint = 0

	for current := l.Head; current != nil; current = current.Next {
		size++
	}

	return size
}

func (l LinkedList) Get(position uint) *Node {
	current := l.Head
	var i uint
	for i = 0; i < position; i++ {
		if current == nil {
			return nil
		}

		current = current.Next
	}

	return current
}

func (l *LinkedList) Insert(position uint, value int) {
	if position > l.Size() {
		return
	}

	if l.Head == nil || position == 0 {
		l.Head = &Node{Next: l.Head, Value: value}
		return
	}

	previousNode := l.Get(position - 1)
	if previousNode == nil {

	}
	nextNode := previousNode.Next

	newNode := &Node{Next: nextNode, Value: value}
	previousNode.Next = newNode
}

func (l *LinkedList) Remove(position uint) *Node {
	if position > l.Size() {
		return nil
	}

	if position == 0 {
		targetNode := l.Head
		l.Head = l.Head.Next
		return targetNode
	}

	previousNode := l.Get(position - 1)
	targetNode := previousNode.Next
	previousNode.Next = targetNode.Next

	return targetNode
}

func (l LinkedList) Print() {
	fmt.Print("[")

	if l.Head != nil {
		for node := l.Head; node != nil; node = node.Next {
			fmt.Printf("%v", node.Value)

			if node.Next != nil {
				fmt.Printf(" ")
			}
		}
	}

	fmt.Print("]")
}
