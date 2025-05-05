package list

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	len  int
	head *Node
	tail *Node
}

type Node struct {
	val  any
	prev *Node
	next *Node
}

func NewList() *LinkedList {
	return &LinkedList{
		len: 0,
	}
}

func NewNode(val any) *Node {
	return &Node{
		val: val,
	}
}

func (l *LinkedList) Prepend(val any) {
	node := NewNode(val)

	if l.len <= 0 {
		l.head = node
		l.tail = node
	} else {
		exHead := l.head
		exHead.prev = node
		node.next = l.head
		l.head = node
	}

	l.len++
}

func (l *LinkedList) Append(val any) {
	node := NewNode(val)

	if l.len <= 0 {
		l.head = node
		l.tail = node
	} else {
		exTail := l.tail
		exTail.next = node
		node.prev = exTail
		l.tail = node
	}

	l.len++
}

func (l *LinkedList) AppendToIndex(val any, index int) error {
	if index > l.len {
		return errors.New("Index out of bounds.")
	}

	if index == 0 || l.len == 0 {
		l.Prepend(val)
		return nil
	}

	if index == l.len-1 {
		l.Append(val)
		return nil
	}

	node := NewNode(val)

	current, err := l.Get(index)
	if err != nil {
		return err
	}

	current.prev.next = node
	node.prev = current.prev.next

	current.prev = node
	node.next = current

	l.len++

	return nil
}

func (l *LinkedList) IsEmpty() bool {
	return l.len <= 0
}

func (l *LinkedList) Get(index int) (*Node, error) {
	if index >= l.len || index < 0 {
		return nil, errors.New("Index out of bounds.")
	}

	lap := 0
	current := l.head

	for index > lap {
		current = current.next
		lap++
	}

	return current, nil
}

func (l *LinkedList) RemoveAtIndex(index int) (*Node, error) {
	if index >= l.len {
		return &Node{}, errors.New("Index out of bounds.")
	}

	current, err := l.Get(index)
	if err != nil {
		return &Node{}, err
	}

	if current.prev == nil {
		l.head = current.next
		l.len--
		return current, nil
	}

	if current.next == nil {
		l.tail = current.prev
		l.len--
		return current, nil
	}

	current.prev.next = current.next
	l.len--
	return current, nil
}

func (l *LinkedList) Find(val any) (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("List is empty.")
	}

	index := 0
	current := l.head
	for current != nil {
		if current.val == val {
			return index, nil
		}

		current = current.next
		index++
	}

	return 0, errors.New("Value not found.")
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *LinkedList) String() string {
	result := ""
	current := l.head
	for current != nil {
		result += fmt.Sprintf("%v <-> ", current.val)
		current = current.next
	}
	result += "nil"
	return result
}
