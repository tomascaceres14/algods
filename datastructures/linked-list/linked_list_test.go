package list

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewList()

	if !list.IsEmpty() || list.head != nil || list.tail != nil {
		t.Error("List should be initialized empty")
	}
}

func NewDefaultList() *LinkedList {
	list := NewList()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	return list
}

func TestAppend(t *testing.T) {
	list := NewDefaultList()

	test := &Node{
		val: 10,
		next: &Node{
			val: 20,
		},
	}

	if list.head.next != list.tail.prev {
		t.Error("Expected head.next and tail.prev to be equal, got head.next", list.head.next, "and head.prev", list.head.prev)
	}

	if list.head.val != test.val {
		t.Error("Expected list.head.next.val and test.next.val to be equal, got", list.head.next.val, "and test.next.val", test.next.val)
	}
}

func TestPrepend(t *testing.T) {
	list := NewList()

	list.Prepend(10)
	list.Prepend(20)
	list.Prepend(30)

	test := &Node{
		val: 30,
		next: &Node{
			val: 20,
		},
	}

	if list.head.next != list.tail.prev {
		t.Error("Expected head.next and tail.prev to be equal, got head.next", list.head.next, "and head.prev", list.head.prev)
	}

	if list.head.next.val != test.next.val {
		t.Error("Expected list.head.next.val and test.next.val to be equal, got", list.head.next.val, "and test.next.val", test.next.val)
	}

	if list.head.val != test.val {
		t.Error("Expected list.head.next.val and test.next.val to be equal, got", list.head.next.val, "and test.next.val", test.next.val)
	}
}

func TestSearch(t *testing.T) {
	list := NewDefaultList()

	searchIndex := 1

	getNodeResult, err := list.Get(searchIndex)
	if err != nil {
		t.Error(err)
	}

	findNodeResult, err := list.Find(getNodeResult.val)
	if err != nil {
		t.Error(err)
	}

	if getNodeResult.val != 20 {
		t.Errorf("Expected 20, got %v", getNodeResult.val)
	}

	if findNodeResult != searchIndex {
		t.Errorf("Expected %v, got %v", searchIndex, getNodeResult.val)
	}
}

func TestAppendToIndex(t *testing.T) {
	list := NewDefaultList()

	if err := list.AppendToIndex(15, 1); err != nil {
		t.Error(err)
	}

	if err := list.AppendToIndex(5, 0); err != nil {
		t.Error(err)
	}

	if err := list.AppendToIndex(25, list.len-1); err != nil {
		t.Error(err)
	}

	if err := list.AppendToIndex(35, list.len); err == nil {
		t.Error(err)
	}

	fmt.Println(list)

}
