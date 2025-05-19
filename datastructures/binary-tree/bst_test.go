package bst

import (
	"fmt"
	"testing"
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
	tree := NewBTree()

	nums := []int{1, 4, 6, 9, 11, 15, 20, 21, 26, 30}

	for _, v := range nums {
		tree.Insert(v)
	}

	fmt.Println(tree)
}
func TestInsertNode(t *testing.T) {
	tree := NewBTree()

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
