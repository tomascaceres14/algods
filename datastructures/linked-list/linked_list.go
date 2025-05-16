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
	Val  any
	prev *Node
	next *Node
}

func NewList() *LinkedList {
	return &LinkedList{
		len: 0,
	}
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Prepend(val any) {
	node := &Node{
		Val: val,
	}

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
	node := &Node{
		Val: val,
	}

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

func (l *LinkedList) InsertAt(val any, index int) error {
	if index > l.len {
		return errors.New("Index out of bounds.")
	}

	if index == 0 || l.len == 0 {
		l.Prepend(val)
		return nil
	}

	if index == l.len {
		l.Append(val)
		return nil
	}

	node := &Node{
		Val: val,
	}

	current := l.Get(index)

	current.prev.next = node
	node.prev = current.prev

	current.prev = node
	node.next = current

	l.len++

	return nil
}

func (l *LinkedList) IsEmpty() bool {
	return l.len <= 0
}

func (l *LinkedList) Get(index int) *Node {
	if index >= l.len || index < 0 {
		return nil
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current
}

func (l *LinkedList) Remove(val any) error {
	if l.Len() == 0 {
		return errors.New("List is empty.")
	}

	if l.head.Val == val {
		l.head = l.head.next
		l.len--
		return nil
	}

	founded := 0
	l.ForEach(func(index int, node *Node) {
		if node.Val == val && founded == 0 {
			node.prev.next, node.next.prev = node.next, node.prev
			l.len--
			founded = 1
		}
	})

	return nil
}

func (l *LinkedList) RemoveAt(index int) error {
	if index >= l.len {
		return errors.New("Index out of bounds.")
	}

	current := l.Get(index)
	if current == nil {
		return errors.New("Index out of bounds.")
	}

	if current.prev == nil {
		l.head = current.next
		if l.head != nil {
			l.head.prev = nil
		}
		l.len--
		return nil
	}

	if current.next == nil {
		l.tail = current.prev
		if l.tail != nil {
			l.tail.next = nil
		}
		l.len--
		return nil
	}

	current.prev.next = current.next
	l.len--
	return nil
}

func (l *LinkedList) Lsearch(val any) int {
	if l.IsEmpty() {
		return -1
	}

	index := 0
	current := l.head
	for current != nil {
		if current.Val == val {
			return index
		}

		current = current.next
		index++
	}

	return -1
}

func (l *LinkedList) Flat() []any {
	var list []any = []any{}

	if l.len == 0 {
		return list
	}

	current := l.head
	for i := 0; i < l.len; i++ {
		list = append(list, current.Val)
		current = current.next
	}

	return list
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
		result += fmt.Sprintf("%v<->", current.Val)
		current = current.next
	}
	result += "nil"
	return result
}

func (l *LinkedList) ForEach(f func(index int, node *Node)) {
	index := 0
	for node := l.head; node != nil; node = node.next {
		f(index, node)
		index++
	}
}
