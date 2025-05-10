package ht

import list "github.com/tomascaceres14/algods/datastructures/linked-list"

type HashTable struct {
	table        []*list.LinkedList
	size, bucket int
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
		table:  table,
		size:   0,
		bucket: cap,
	}
}

func (ht *HashTable) rollingHash(s string) int64 {
	const m int64 = 1_000_000_009

	var hash int64 = 0
	for i := 0; i < len(s); i++ {
		hash = ((hash<<5 - hash) + int64(s[i])) % m
	}
	return hash
}

func (ht *HashTable) Put(key string, value any) {
	index := ht.rollingHash(key)

	item := item{
		key:   key,
		value: value,
	}

	list := ht.table[index]

	if exists := list.Lsearch(item); exists == -1 {
		list.Append(item)
	} else {
		list.Get(exists).Val = item
	}
}
