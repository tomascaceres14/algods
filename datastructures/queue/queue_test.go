package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := New()

	if queue.len != 0 || !queue.IsEmpty() || queue.items == nil {
		t.Error("Error initializing queue.")
	}
}

func TestPushAndPeek(t *testing.T) {
	queue := New()

	queue.Push("hola")
	queue.Push("chau")

	peek, err := queue.Peek()
	if err != nil {
		t.Error("Error peeking.")
	}

	if peek != "chau" {
		t.Errorf("Push() expected 'chau', got %s.", peek)
	}
}

func TestShift(t *testing.T) {
	queue := New()

	queue.Push("hola")
	queue.Push("chau")

	shift, err := queue.Shift()
	if err != nil {
		t.Error("Error removing last element.")
	}

	if shift != "chau" {
		t.Errorf("Shift() expected 'chau', got %s.", shift)
	}

	if queue.len != 1 {
		t.Errorf("Shift() expected queue.len == 1, got %v.", queue.len)
	}
}

func TestQueue(t *testing.T) {
	queue := New()
	itemsToAdd := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, v := range itemsToAdd {
		queue.Push(v)
	}

	sumPairs := 0
	for queue.Size() > 0 {
		item, _ := queue.Shift()

		num := item.(int)

		if num%2 == 0 {
			sumPairs += num
		}
	}

	if sumPairs != 20 {
		t.Errorf("Expecting 20, got %v", sumPairs)
	}
}
