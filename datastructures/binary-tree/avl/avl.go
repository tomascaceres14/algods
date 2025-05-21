package avl

type Node struct {
	Val                 int
	left, right, parent *Node
	height              int
}

type AVLTree struct {
	root *Node
	size int
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		size: 0,
	}
}

func rotateLeft(node *Node) *Node {

}

func (t *AVLTree) Insert(val int) (*Node, bool) {

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
				return node, true
			}

		} else {
			if current.left != nil {
				current = current.left
			} else {
				node.parent = current
				current.left = node
				t.size++
				return node, true
			}
		}
	}
	return nil, false
}
