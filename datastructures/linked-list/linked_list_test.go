package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_AppendAndPrepend(t *testing.T) {
	ll := NewList()

	ll.Append("A")
	ll.Prepend("B")
	ll.Append("C")

	assert.Equal(t, 3, ll.Len())
	assert.Equal(t, "B", ll.Get(0).Val)
	assert.Equal(t, "A", ll.Get(1).Val)
	assert.Equal(t, "C", ll.Get(2).Val)
}

func TestLinkedList_Get_OutOfBounds(t *testing.T) {
	ll := NewList()
	assert.Nil(t, ll.Get(-1))
	assert.Nil(t, ll.Get(0))

	ll.Append("X")
	assert.Nil(t, ll.Get(1))
}

func TestLinkedList_AppendToIndex(t *testing.T) {
	ll := NewList()
	err := ll.AppendToIndex("A", 0)
	assert.Nil(t, err)
	err = ll.AppendToIndex("B", 1)
	assert.Nil(t, err)
	err = ll.AppendToIndex("C", 1)
	assert.Nil(t, err)

	assert.Equal(t, "A", ll.Get(0).Val)
	assert.Equal(t, "C", ll.Get(1).Val)
	assert.Equal(t, "B", ll.Get(2).Val)
}

func TestLinkedList_AppendToIndex_OutOfBounds(t *testing.T) {
	ll := NewList()
	err := ll.AppendToIndex("A", 2)
	assert.NotNil(t, err)
}

func TestLinkedList_RemoveAtIndex(t *testing.T) {
	ll := NewList()
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	node, err := ll.RemoveAtIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, "B", node.Val)
	assert.Equal(t, 2, ll.Len())

	assert.Equal(t, "A", ll.Get(0).Val)
	assert.Equal(t, "C", ll.Get(1).Val)
}

func TestLinkedList_RemoveAtIndex_FirstAndLast(t *testing.T) {
	ll := NewList()
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	_, err := ll.RemoveAtIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, "B", ll.Get(0).Val)

	_, err = ll.RemoveAtIndex(1)
	assert.Nil(t, err)
	assert.Equal(t, "B", ll.Get(0).Val)
	assert.Nil(t, ll.Get(1))
}

func TestLinkedList_IsEmpty(t *testing.T) {
	ll := NewList()
	assert.True(t, ll.IsEmpty())
	ll.Append("A")
	assert.False(t, ll.IsEmpty())
}

func TestLinkedList_Lsearch(t *testing.T) {
	ll := NewList()
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	assert.Equal(t, 1, ll.Lsearch("B"))
	assert.Equal(t, -1, ll.Lsearch("X"))
}

func TestLinkedList_Clear(t *testing.T) {
	ll := NewList()
	ll.Append("A")
	ll.Append("B")

	ll.Clear()
	assert.True(t, ll.IsEmpty())
	assert.Nil(t, ll.Get(0))
}
