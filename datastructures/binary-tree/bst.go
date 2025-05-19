package bst

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

type BTree struct {
	root *Node
	size int
}

type Node struct {
	Val                 int
	parent, left, right *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

func NewBTree() *BTree {
	return &BTree{
		size: 0,
	}
}

func NewBTreeFromArray(array []int) *BTree {

	bst := &BTree{}

	for _, v := range array {
		bst.Insert(v)
	}

	return bst
}

func (t *BTree) Len() int {
	return t.size
}

func (t *BTree) Insert(val int) error {
	node := NewNode(val)

	if t.root == nil {
		t.root = node
	}

	current := t.root
	for {
		if current.Val == node.Val {
			return errors.New("Value " + strconv.Itoa(node.Val) + " already in tree.")
		}

		if current.Val < node.Val {
			if current.right != nil {
				current = current.right
			} else {
				node.parent = current
				current.right = node
				t.size++
				return nil
			}

		} else {
			if current.left != nil {
				current = current.left
			} else {
				node.parent = current
				current.left = node
				t.size++
				return nil
			}
		}
	}
}

func (t *BTree) InsertMany(many ...int) error {
	for _, v := range many {
		if err := t.Insert(v); err != nil {
			return err
		}
	}

	return nil
}

func (t *BTree) Min() *Node {
	if t.root == nil {
		return &Node{}
	}

	current := t.root
	for current.left != nil {
		current = current.left
	}

	return current
}

func (t *BTree) InOrder() []int {
	var result []int
	inOrderHelper(t.root, &result)
	return result
}

func (t *BTree) Balance() {
	inorder := t.InOrder()

	bst := NewBTree()

	for range inorder {
		mid := len(inorder) / 2
		bst.Insert(inorder[mid])
		inorder = slices.Delete(inorder, mid, mid+1)
	}

	t.root = bst.root
}

func inOrderHelper(n *Node, result *[]int) {
	if n == nil {
		return
	}
	inOrderHelper(n.left, result)
	*result = append(*result, n.Val)
	inOrderHelper(n.right, result)
}

func stringify(n *Node, level int) string {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%d\n", n.Val)
		stringify(n.left, level)

		return format
	} else {
		return ""
	}
}

func (bst *BTree) String() string {
	return stringify(bst.root, 0)
}
