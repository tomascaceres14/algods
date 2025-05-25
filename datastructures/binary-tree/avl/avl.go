package avl

import (
	"errors"
	"fmt"

	"github.com/tomascaceres14/algods/datastructures/queue"
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Val         T
	left, right *Node[T]
	height      int
}

type AVLTree[T constraints.Ordered] struct {
	root *Node[T]
	size int
}

func NewNode[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{
		Val:    val,
		height: 1,
	}
}

func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		size: 0,
	}
}

func (t *AVLTree[T]) Insert(val T) {
	t.root = insert(val, t.root)
}

func insert[T constraints.Ordered](val T, node *Node[T]) *Node[T] {
	if node == nil {
		return NewNode(val)
	}

	if val < node.Val {
		node.left = insert(val, node.left)
	} else if val > node.Val {
		node.right = insert(val, node.right)
	} else {
		return node
	}

	updateHeight(node)
	return applyRotation(node)
}

func (t *AVLTree[T]) Delete(val T) {
	t.root = delete(val, t.root)
}

func delete[T constraints.Ordered](val T, node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	if val < node.Val {
		node.left = delete(val, node.left)
	} else if val > node.Val {
		node.right = delete(val, node.right)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		nextInOrder, _ := getMin(node.right)
		node.Val = nextInOrder
		node.right = delete(node.Val, node.right)
	}

	updateHeight(node)
	return applyRotation(node)
}

func (t *AVLTree[T]) Exists(val T) bool {
	if t.root == nil {
		return false
	}

	return exists(val, t.root)
}

func exists[T constraints.Ordered](val T, node *Node[T]) bool {
	if node == nil {
		return false
	}

	if val < node.Val {
		return exists(val, node.left)
	} else if val > node.Val {
		return exists(val, node.right)
	}

	return true
}

func (t *AVLTree[T]) InOrder() {
	inOrder(t.root)
	fmt.Println("")
}

func inOrder[T constraints.Ordered](node *Node[T]) {
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Printf("%v ", node.Val)
	inOrder(node.right)
}

func (t *AVLTree[T]) PreOrder() {
	preOrder(t.root)
	fmt.Println("")
}

func preOrder[T constraints.Ordered](node *Node[T]) {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.Val)
	preOrder(node.left)
	preOrder(node.right)
}

func (t *AVLTree[T]) PostOrder() {
	postOrder(t.root)
	fmt.Println("")
}

func postOrder[T constraints.Ordered](node *Node[T]) {
	if node == nil {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Printf("%v ", node.Val)
}

func (t *AVLTree[T]) LevelOrder() {
	queue := queue.New()
	level := 1
	queue.Push(t.root)
	fmt.Printf("%v\n", t.root.Val)

	for !queue.IsEmpty() {
		curr, _ := queue.Shift()
		node := curr.(*Node[T])

		if node.left != nil {
			fmt.Printf("\n%v - ", node.left.Val)
			queue.Push(node.left)
		}

		if node.right != nil {
			fmt.Printf("%v\n", node.right.Val)
			queue.Push(node.right)
		}

		level++
	}
}

func applyRotation[T constraints.Ordered](node *Node[T]) *Node[T] {
	balance := getBalance(node)

	if balance > 1 {
		println("--- WAR: Tree unbalanced. Rotating right")
		if getBalance(node.left) < 0 {
			node.left = rotateLeft(node.left)
		}
		return rotateRight(node)
	}

	if balance < -1 {
		println("--- WAR: Tree unbalanced. Rotating left")
		if getBalance(node.right) > 0 {
			node.right = rotateRight(node.right)
		}
		return rotateLeft(node)
	}

	return node
}

func rotateLeft[T constraints.Ordered](x *Node[T]) *Node[T] {
	if x == nil {
		return x
	}

	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

func rotateRight[T constraints.Ordered](x *Node[T]) *Node[T] {
	if x == nil {
		return x
	}

	y := x.left
	T2 := y.right

	y.right = x
	x.left = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

func updateHeight[T constraints.Ordered](node *Node[T]) {
	node.height = max(getHeight(node.left), getHeight(node.right)) + 1
}

func getBalance[T constraints.Ordered](node *Node[T]) int {
	if node == nil {
		return 0
	}

	balance := getHeight(node.left) - getHeight(node.right)
	return balance
}

func getHeight[T constraints.Ordered](node *Node[T]) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (t *AVLTree[T]) Min() (T, error) {
	var nilVal T
	if min, err := getMin(t.root); err != nil {
		return nilVal, err
	} else {
		return min, nil
	}
}

func getMin[T constraints.Ordered](node *Node[T]) (T, error) {

	var nilVal T

	if node == nil {
		return nilVal, errors.New("node is empty")
	}

	if node.left == nil {
		return node.Val, nil
	}

	nilVal, _ = getMin(node.left)

	return nilVal, nil
}

func (t *AVLTree[T]) Max() (T, error) {
	var nilVal T
	if max, err := getMax(t.root); err != nil {
		return nilVal, err
	} else {
		return max, nil
	}
}

func getMax[T constraints.Ordered](node *Node[T]) (T, error) {

	var nilVal T

	if node == nil {
		return nilVal, errors.New("node is empty")
	}

	if node.right == nil {
		return node.Val, nil
	}

	nilVal, _ = getMax(node.right)

	return nilVal, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func stringify[T constraints.Ordered](n *Node[T], level int) string {
	if n != nil {
		format := ""
		for range level {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%v\n", n.Val)
		stringify(n.left, level)

		return format
	} else {
		return ""
	}
}

func (bst *AVLTree[T]) String() string {
	return stringify(bst.root, 0)
}
