package linkedlist

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	ll := &LinkedList{}

	if !ll.IsEmpty() {
		t.FailNow()
	}
}

func TestIsNotEmpty(t *testing.T) {
	ll := &LinkedList{}

	ll.Insert(0, 10)

	if ll.IsEmpty() {
		t.FailNow()
	}
}

func TestEmptySize(t *testing.T) {
	ll := &LinkedList{}

	if ll.Size() != 0 {
		t.FailNow()
	}
}

func TestSize(t *testing.T) {
	ll := &LinkedList{}

	ll.Insert(0, 10)
	ll.Insert(1, 11)

	if ll.Size() != 2 {
		t.FailNow()
	}
}

func TestEmptyGet(t *testing.T) {
	ll := &LinkedList{}

	node := ll.Get(0)

	if node != nil {
		t.FailNow()
	}
}

func TestOutOfRangeGet(t *testing.T) {
	ll := &LinkedList{}

	ll.Insert(0, 10)
	ll.Insert(1, 11)
	ll.Insert(2, 22)

	nodeOK := ll.Get(2)
	nodeNil := ll.Get(100)

	if nodeOK.Value == 22 && nodeNil != nil {
		t.FailNow()
	}
}

func TestInsertAndGet(t *testing.T) {
	ll := &LinkedList{}

	ll.Insert(0, 10)
	node := ll.Get(0)

	if node.Value != 10 {
		t.FailNow()
	}
}

func TestInsertAndGet2(t *testing.T) {
	ll := &LinkedList{}

	ll.Insert(0, 10)
	ll.Insert(1, 11)
	ll.Insert(2, 22)

	node1 := ll.Get(1)
	node2 := ll.Get(2)

	if node1.Value != 11 && node2.Value != 22 {
		t.FailNow()
	}
}

func TestInsertAndGet3(t *testing.T) {
	ll := &LinkedList{}

	ll.Insert(0, 10)
	ll.Insert(0, 9)

	node0 := ll.Get(0)
	node1 := ll.Get(1)

	if node0.Value != 9 && node1.Value != 10 {
		t.FailNow()
	}
}

func TestInsertInvalidPosition(t *testing.T) {
	ll := &LinkedList{}

	ll.Insert(100, 10)

	if ll.Size() != 0 {
		t.FailNow()
	}
}

func TestEmptyRemove(t *testing.T) {
	ll := &LinkedList{}

	nodeNil := ll.Remove(100)

	if nodeNil != nil {
		t.FailNow()
	}
}

func TestRemove1(t *testing.T) {
	ll := &LinkedList{}

	ll.Insert(0, 10)
	ll.Insert(1, 11)
	ll.Insert(2, 22)

	removedNode11 := ll.Remove(1)
	removedNode10 := ll.Remove(0)
	node22 := ll.Get(0)
	size := ll.Size()

	if removedNode11.Value != 11 && removedNode10.Value != 10 && node22.Value != 22 && size != 1 {
		t.FailNow()
	}
}

func TestRemove2(t *testing.T) {
	ll := &LinkedList{}

	nodeNil := ll.Remove(100)

	if nodeNil != nil {
		t.FailNow()
	}
}

func TestRemove3(t *testing.T) {
	ll := &LinkedList{}

	if nodeNil := ll.Remove(0); nodeNil != nil {
		t.FailNow()
	}
}

func BenchmarkInsertAndRemoveMany(b *testing.B) {
	const nodesToInsert = 10000

	for i := 0; i < b.N; i++ {
		ll := &LinkedList{}

		for n := 0; n < nodesToInsert; n++ {
			ll.Insert(0, n)
		}

		for n := 0; n < nodesToInsert; n++ {
			ll.Remove(0)
		}
	}
}

func BenchmarkSize(b *testing.B) {
	const nodesToTraverse = 1000

	ll := &LinkedList{}
	for n := 0; n < nodesToTraverse; n++ {
		ll.Insert(0, n)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.Size()
	}
}

func ExamplePrintEmptyList() {
	ll := &LinkedList{}

	ll.Print()

	// Output: []
}

func ExamplePrintList() {
	ll := &LinkedList{}

	ll.Insert(0, 2)
	ll.Insert(0, 1)
	ll.Insert(0, 0)

	ll.Print()

	// Output: [0 1 2]
}
