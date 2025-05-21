package bst

import (
	"errors"
	"fmt"
	"slices"
)

type BSTree struct {
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

func NewBSTree() *BSTree {
	return &BSTree{
		size: 0,
	}
}

func NewBSTreeFromArray(array []int) *BSTree {

	bst := &BSTree{}

	for _, v := range array {
		bst.Insert(v)
	}

	return bst
}

func (t *BSTree) Len() int {
	return t.size
}

func (t *BSTree) Insert(val int) (*Node, bool) {
	node := NewNode(val)

	if t.root == nil {
		t.root = node
		return node, true
	}

	current := t.root
	for {
		if current.Val == node.Val {
			return nil, false
		}

		if current.Val < node.Val {
			if current.right != nil {
				current = current.right
			} else {
				node.parent = current
				current.right = node
				t.size++
			}

		} else {
			if current.left != nil {
				current = current.left
			} else {
				node.parent = current
				current.left = node
				t.size++
			}
		}
	}

}

func (t *BSTree) Search(val int) (*Node, error) {
	current := t.root

	for current != nil {
		if current.Val == val {
			return current, nil
		}

		if val < current.Val {
			current = current.left
		} else {
			current = current.right
		}
	}

	return &Node{}, errors.New("Node not found")
}

func (t *BSTree) Delete(val int) (*Node, error) {
	var parent *Node
	current := t.root
	for current != nil {
		if current.Val == val {
			break
		} else if val < current.Val {
			parent = current
			current = current.left
		} else {
			parent = current
			current = current.right
		}
	}

	if current == nil {
		return nil, errors.New("Node not found.")
	}

	result := current

	defer func() {
		t.size--
	}()

	// Leaf deletion
	if current.left == nil && current.right == nil {
		if t.root == parent {
			t.root = nil
		} else if parent.left == current {
			parent.left = nil
		} else {
			parent.right = nil
		}
		t.size--
		return result, nil
	}

	// Only right child
	if current.left == nil {
		current.right.parent = current.parent
		if parent.left == current {
			parent.left = current.right
		} else {
			parent.right = current.right
		}
		return result, nil
	}

	// Only left child
	if current.right == nil {
		current.left.parent = current.parent
		if parent.right == current {
			parent.right = current.left
		} else {
			parent.left = current.left
		}
		return result, nil
	}

	// Having both childs, find next in-order
	successor := current.right
	for successor.left != nil {
		successor = successor.left
	}

	// Switch values
	current.Val = successor.Val

	// find successor's parent and child for validations
	succParent := successor.parent
	child := successor.right

	if succParent.left == successor {
		succParent.left = child
	} else {
		succParent.right = child
	}

	if child != nil {
		child.parent = succParent
	}

	return result, nil
}

func inOrderHelper(n *Node, result *[]int) {
	if n == nil {
		return
	}
	inOrderHelper(n.left, result)
	*result = append(*result, n.Val)
	inOrderHelper(n.right, result)
}

func (t *BSTree) InOrder() []int {
	var result []int
	inOrderHelper(t.root, &result)
	return result
}

func (t *BSTree) Balance() {
	inorder := t.InOrder()

	bst := NewBSTree()

	for range inorder {
		mid := len(inorder) / 2
		bst.Insert(inorder[mid])
		inorder = slices.Delete(inorder, mid, mid+1)
	}

	t.root = bst.root
}

func stringify(n *Node, level int) string {
	if n != nil {
		format := ""
		for range level {
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

func (bst *BSTree) String() string {
	return stringify(bst.root, 0)
}
