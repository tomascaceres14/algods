package list

import (
	"testing"
)

func TestLinkedListBasicOps(t *testing.T) {
	l := NewList()

	if !l.IsEmpty() {
		t.Error("Expected list to be empty")
	}

	l.Append("a")
	l.Append("b")
	l.Prepend("c")

	if l.Len() != 3 {
		t.Errorf("Expected length 3, got %d", l.Len())
	}

	if l.head.Val != "c" || l.tail.Val != "b" {
		t.Error("Head or tail values are incorrect after appends and prepends")
	}

	expected := "c <-> a <-> b <-> nil"
	if l.String() != expected {
		t.Errorf("Expected string %q, got %q", expected, l.String())
	}
}

func TestLinkedListGet(t *testing.T) {
	l := NewList()
	l.Append("x")
	l.Append("y")
	l.Append("z")

	if node := l.Get(1); node == nil || node.Val != "y" {
		t.Error("Get(1) failed to retrieve correct node")
	}

	if node := l.Get(3); node != nil {
		t.Error("Expected Get(3) to return nil (out of bounds)")
	}
}

func TestAppendToIndex(t *testing.T) {
	l := NewList()
	l.Append("1")
	l.Append("2")
	l.Append("4")
	err := l.AppendToIndex("3", 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if l.Get(2).Val != "3" {
		t.Errorf("Expected value at index 2 to be '3', got %v", l.Get(2).Val)
	}

	err = l.AppendToIndex("fail", 99)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestRemoveAtIndex(t *testing.T) {
	l := NewList()
	l.Append("a")
	l.Append("b")
	l.Append("c")

	node, err := l.RemoveAtIndex(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if node.Val != "b" {
		t.Errorf("Expected to remove 'b', got %v", node.Val)
	}
	if l.Len() != 2 {
		t.Errorf("Expected length 2 after removal, got %d", l.Len())
	}

	_, err = l.RemoveAtIndex(10)
	if err == nil {
		t.Error("Expected error when removing at invalid index")
	}
}

func TestLsearch(t *testing.T) {
	l := NewList()
	l.Append("apple")
	l.Append("banana")
	l.Append("cherry")

	if idx := l.Lsearch("banana"); idx != 1 {
		t.Errorf("Expected index 1 for 'banana', got %d", idx)
	}

	if idx := l.Lsearch("missing"); idx != -1 {
		t.Errorf("Expected -1 for 'missing', got %d", idx)
	}
}

func TestClear(t *testing.T) {
	l := NewList()
	l.Append("1")
	l.Append("2")
	l.Clear()

	if l.Len() != 0 || !l.IsEmpty() {
		t.Error("Clear() did not reset the list properly")
	}
	if l.head != nil || l.tail != nil {
		t.Error("Head or tail should be nil after Clear()")
	}
}
