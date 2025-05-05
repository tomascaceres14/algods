package queue

import "errors"

type Queue struct {
	items []any
	len   int
}

func New() *Queue {
	return &Queue{
		items: []any{},
		len:   0,
	}
}

func (q *Queue) Push(item any) {
	q.items = append(q.items, item)
	q.len++
}

func (q *Queue) Shift() (any, error) {

	if q.len <= 0 {
		return nil, errors.New("queue is empty")
	}

	head := q.items[0]
	q.items = q.items[1:]
	q.len--

	return head, nil
}

func (q *Queue) Peek() (any, error) {
	if q.len <= 0 {
		return nil, errors.New("queue is empty")
	}

	return q.items[0], nil
}

func (q *Queue) Size() int {
	return q.len
}

func (q *Queue) IsEmpty() bool {
	return q.len <= 0
}
