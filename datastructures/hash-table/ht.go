package ht

import (
	"fmt"

	list "github.com/tomascaceres14/algods/datastructures/linked-list"
)

type HashTable struct {
	Table            []*list.LinkedList
	size, buckets    int
	minSize, maxSize float64
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

	minSize := 0.25 * float64(cap)
	maxSize := 0.65 * float64(cap)

	return &HashTable{
		Table:   table,
		size:    0,
		buckets: cap,
		minSize: minSize,
		maxSize: maxSize,
	}
}

func (ht *HashTable) Size() int {
	return ht.size
}

func (ht *HashTable) Buckets() int {
	return ht.buckets
}

func (ht *HashTable) Put(key string, value any) {
	if ht.size >= int(ht.maxSize) {
		ht.expand()
	}

	index := ht.rollingHash(key)
	list := ht.Table[index]

	if exists, ok := findItemByKey(list, key); !ok {
		list.Append(&item{
			key:   key,
			value: value,
		})
	} else {
		exists.value = value
	}

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

func (ht *HashTable) Delete(key string) {
	if ht.size <= int(ht.minSize) {
		ht.shrink()
	}

	index := ht.rollingHash(key)
	bucket := ht.Table[index]

	if bucket.Len() == 0 {
		return
	}

	if bucket.Len() == 1 {
		bucket.RemoveAt(0)
		ht.size--
		return
	}

	bucket.ForEach(func(index int, node *list.Node) {
		item := node.Val.(*item)
		if item.key == key {
			bucket.RemoveAt(index)
			ht.size--
			return
		}
	})
}

func (ht *HashTable) rollingHash(s string) int64 {
	const m int64 = 1_000_000_009

	var hash int64 = 0
	for i := 0; i < len(s); i++ {
		hash = ((hash<<5 - hash) + int64(s[i])) % m
	}
	return hash % int64(ht.buckets)
}

func (ht *HashTable) expand() {
	newTable := New(ht.buckets * 2)
	ht.ForEach(func(i *item) {
		newTable.Put(i.key, i.value)
	})

	ht.Table = newTable.Table
	ht.buckets = newTable.buckets
	ht.maxSize = newTable.maxSize
	ht.minSize = newTable.minSize
}

func (ht *HashTable) shrink() {
	newTable := New(int(ht.buckets / 2))
	ht.ForEach(func(i *item) {
		newTable.Put(i.key, i.value)
	})

	ht.Table = newTable.Table
	ht.buckets = newTable.buckets
	ht.maxSize = newTable.maxSize
	ht.minSize = newTable.minSize
}

func (ht *HashTable) ForEach(f func(*item)) {
	for _, bucket := range ht.Table {
		if !bucket.IsEmpty() {
			bucket.ForEach(func(index int, node *list.Node) {
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
