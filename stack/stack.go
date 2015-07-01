package stack

import (
	"github.com/jeffbmartinez/datastructures/linkedlist"
)

type Stack interface {
	IsEmpty() bool
	Push(value int)
	Pop() (int, error)
	Size() uint
}

type LinkedStack struct {
	linkedlist.LinkedList
}

func (s LinkedStack) IsEmpty() bool {
	return s.LinkedList.IsEmpty()
}

func (s *LinkedStack) Push(value int) {
	s.Insert(0, value)
}

func (s *LinkedStack) Pop() (value int, ok bool) {
	node := s.Remove(0)

	if node == nil {
		return 0, false
	}

	return node.Value, true
}

func (s LinkedStack) Size() uint {
	return s.LinkedList.Size()
}
