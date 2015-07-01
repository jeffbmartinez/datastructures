package queue

import (
	"github.com/jeffbmartinez/datastructures/linkedlist"
)

type Stack interface {
	IsEmpty() bool
	Enqueue(value int)
	Dequeue() (int, error)
	Size() uint
}

type LinkedQueue struct {
	linkedlist.LinkedList
}

func (q LinkedQueue) IsEmpty() bool {
	return q.LinkedList.IsEmpty()
}

func (q *LinkedQueue) Enqueue(value int) {
	q.Insert(q.Size(), value)
}

func (q *LinkedQueue) Dequeue() (value int, ok bool) {
	node := q.Remove(0)

	if node == nil {
		return 0, false
	}

	return node.Value, true
}

func (q LinkedQueue) Size() uint {
	return q.LinkedList.Size()
}
