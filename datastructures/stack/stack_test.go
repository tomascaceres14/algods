package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := New()

	if stack.len != 0 || !stack.IsEmpty() || stack.items == nil {
		t.Error("Error initializing stack.")
	}
}

func TestPushAndPeek(t *testing.T) {
	stack := New()

	stack.Push("hola")
	stack.Push("chau")

	peek, err := stack.Peek()
	if err != nil {
		t.Error("Error peeking.")
	}

	if peek != "chau" {
		t.Errorf("Push() expected 'chau', got %s.", peek)
	}
}

func TestPop(t *testing.T) {
	stack := New()

	stack.Push("hola")
	stack.Push("chau")

	pop, err := stack.Pop()
	if err != nil {
		t.Error("Error removing last element.")
	}

	if pop != "chau" {
		t.Errorf("Pop() expected 'chau', got %s.", pop)
	}

	if stack.len != 1 {
		t.Errorf("Pop() expected stack.len == 1, got %v.", stack.len)
	}
}
