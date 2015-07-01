package stack

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	s := &LinkedStack{}

	if !s.IsEmpty() {
		t.FailNow()
	}

	if s.Size() != 0 {
		t.FailNow()
	}
}

func TestPushPop(t *testing.T) {
	s := &LinkedStack{}

	s.Push(2)
	s.Push(1)
	s.Push(0)

	if zero, ok := s.Pop(); !(ok && zero == 0) {
		t.FailNow()
	}

	if one, ok := s.Pop(); !(ok && one == 1) {
		t.FailNow()
	}

	if two, ok := s.Pop(); !(ok && two == 2) {
		t.FailNow()
	}

	if _, ok := s.Pop(); ok {
		t.FailNow()
	}
}

func TestSize(t *testing.T) {
	s := &LinkedStack{}

	s.Push(2)
	s.Push(1)
	s.Push(0)

	if s.Size() != 3 {
		t.FailNow()
	}
}
