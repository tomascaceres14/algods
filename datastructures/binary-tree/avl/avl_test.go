package avl

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	tree := NewAVLTree[string]()

	tree.Insert("h")
	tree.Insert("i")
	tree.Insert("j")
	tree.Insert("k")

	tree.Insert("c")
	tree.Insert("d")
	tree.Insert("a")

	max, err := tree.Max()
	if err != nil {
		t.Error(err)
	}
	min, err := tree.Min()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Current max: ", max)
	fmt.Println("Current min: ", min)

	fmt.Println(tree)
	fmt.Println("--------------------------")
	tree.Delete("a")
	fmt.Println(tree)
}
