package search

import (
	"golang.org/x/exp/constraints"
)

func Bsearch[T constraints.Ordered](arr []T, el T) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (high + low) / 2

		if arr[mid] == el {
			return mid
		}

		if el < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func Lsearch[T constraints.Ordered](arr []T, el T) int {
	for k, v := range arr {
		if v == el {
			return k
		}
	}
	return -1
}
