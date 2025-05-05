package stack

import "errors"

type Stack struct {
	items []any
	len   int
}

func New() *Stack {
	newItems := []any{}
	return &Stack{
		items: newItems,
		len:   0,
	}
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
	s.len++
}

func (s *Stack) Pop() (any, error) {

	if s.len <= 0 {
		return nil, errors.New("stack is empty")
	}

	lastItem := s.items[s.len-1]
	s.items = s.items[:s.len-1]
	s.len--

	return lastItem, nil
}

func (s *Stack) Peek() (any, error) {

	if s.len <= 0 {
		return nil, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) <= 0
}
