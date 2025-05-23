package search

import (
	"fmt"
	"testing"
)

func TestBsearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Expecting 6 got ", Bsearch(arr, 7))
	fmt.Println("Expecting -1 got ", Bsearch(arr, 11))
}

func TestLsearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Expecting 7 got ", Bsearch(arr, 8))
	fmt.Println("Expecting 0 got ", Bsearch(arr, 1))
}
