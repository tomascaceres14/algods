package bst

import (
	"testing"
)

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

	if len(inOrder) != len(sortedNums) {
		t.Error("InOrder() length doesn't match test slice")
	}

	for i, v := range inOrder {
		if v != sortedNums[i] {
			t.Error("Items doesnt match", v, sortedNums[i])
		}
	}
}
