package avl

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	tree := NewAVLTree[string]()

	tree.Insert("e")
	tree.Insert("h")
	tree.Insert("i")
	tree.Insert("b")
	tree.Insert("j")
	tree.Insert("k")
	tree.Insert("f")
	tree.Insert("c")
	tree.Insert("d")
	tree.Insert("g")
	tree.Insert("a")
	tree.Insert("l")
	tree.Insert("m")
	tree.Insert("n")
	tree.Insert("r")
	tree.Insert("o")
	tree.Insert("p")
	tree.Insert("q")
	tree.Insert("s")
	tree.Insert("t")
	tree.Insert("u")
	tree.Insert("v")
	tree.Insert("w")

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
	fmt.Println("--------------------------")
	//tree.Delete("a")

	fmt.Println("---- Level Order ----")
	tree.LevelOrder()
}
