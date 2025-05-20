package bst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func equalSlices(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}

func TestQuickly(t *testing.T) {
	tree := NewBSTree()

	nums := []int{10, 12, 8, 11, 13, 9, 14}

	for _, v := range nums {
		tree.Insert(v)
	}

	searchNode, err := tree.Search(10)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 10, searchNode.Val)

	fmt.Println(tree)
}
func TestInsertNode(t *testing.T) {
	tree := NewBSTree()

	nums := []int{5, 1, 3, 6, 8, 2, 10, 4, 1}
	sortedNums := []int{1, 2, 3, 4, 5, 6, 8, 10}

	for _, v := range nums {
		tree.Insert(v)
	}

	if tree.Min().Val != 1 {
		t.Errorf("Expected min to be 1, got %v", tree.Min().Val)
	}

	inOrder := tree.InOrder()

	if !equalSlices(inOrder, sortedNums) {
		t.Errorf("InOrder() bad order, got %v and expected %v", inOrder, sortedNums)
	}
}
