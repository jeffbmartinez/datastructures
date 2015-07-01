package queue

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	q := &LinkedQueue{}

	if !q.IsEmpty() {
		t.FailNow()
	}

	if q.Size() != 0 {
		t.FailNow()
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := &LinkedQueue{}

	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	if zero, ok := q.Dequeue(); !(ok && zero == 0) {
		t.FailNow()
	}

	if one, ok := q.Dequeue(); !(ok && one == 1) {
		t.FailNow()
	}

	if two, ok := q.Dequeue(); !(ok && two == 2) {
		t.FailNow()
	}

	if _, ok := q.Dequeue(); ok {
		t.FailNow()
	}
}

func TestSize(t *testing.T) {
	q := &LinkedQueue{}

	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size() != 3 {
		t.FailNow()
	}
}
