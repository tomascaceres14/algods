package list

import "testing"

func TestNewLinkedList(t *testing.T) {
	list := NewList()

	if !list.IsEmpty() || list.head != nil || list.tail != nil {
		t.Error("List should be initialized empty")
	}
}

func TestAppend(t *testing.T) {
	list := NewList()
	list.Append(10)
	list.Append(20)
	list.Append(30)

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
