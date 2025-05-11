package ht

import (
	"fmt"

	list "github.com/tomascaceres14/algods/datastructures/linked-list"
)

type HashTable struct {
	Table         []*list.LinkedList
	size, buckets int
}

type item struct {
	key   string
	value any
}

func New(cap int) *HashTable {
	table := make([]*list.LinkedList, cap)
	for i := range cap {
		table[i] = list.NewList()
	}

	return &HashTable{
		Table:   table,
		size:    0,
		buckets: cap,
	}
}

func (ht *HashTable) Size() int {
	return ht.size
}

func (ht *HashTable) Buckets() int {
	return ht.buckets
}

func (ht *HashTable) Put(key string, value any) {
	index := ht.rollingHash(key)

	list := ht.Table[index]

	exists, ok := findItemByKey(list, key)
	if !ok {
		list.Append(&item{
			key:   key,
			value: value,
		})
		return
	}

	exists.value = value
	ht.size++
}

func (ht *HashTable) Get(key string) any {
	index := ht.rollingHash(key)

	list := ht.Table[index]

	if list.Len() == 0 {
		return nil
	}

	item, ok := findItemByKey(list, key)
	if !ok {
		return nil
	}

	return item.value
}

func findItemByKey(list *list.LinkedList, key string) (*item, bool) {
	for i := 0; i < list.Len(); i++ {
		node := list.Get(i)
		item, ok := node.Val.(*item)
		if ok && item.key == key {
			return item, true
		}
	}

	return nil, false
}

func (ht *HashTable) rollingHash(s string) int64 {
	const m int64 = 1_000_000_009

	var hash int64 = 0
	for i := 0; i < len(s); i++ {
		hash = ((hash<<5 - hash) + int64(s[i])) % m
	}
	return hash % int64(ht.buckets)
}

func (ht *HashTable) ForEach(f func(*item)) {
	for _, bucket := range ht.Table {
		if !bucket.IsEmpty() {
			bucket.ForEach(func(node *list.Node) {
				f(node.Val.(*item))
			})
		}
	}
}

func (ht *HashTable) String() string {
	var result string = ""

	result += fmt.Sprintf("Table: %v\n", ht.Table)
	result += fmt.Sprintf("Buckets: %v\n", ht.buckets)
	result += fmt.Sprintf("Size: %v\n", ht.size)

	return result
}

func (it *item) String() string {
	return fmt.Sprintf("{key=%s, val=%v}", it.key, it.value)
}
