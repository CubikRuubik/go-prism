package main

import "github.com/google/uuid"

type QueueItem struct {
	ID    string
	Value int
}

type Queue struct {
	items []QueueItem
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value int) QueueItem {
	item := QueueItem{
		ID:    uuid.New().String(),
		Value: value,
	}
	q.items = append(q.items, item)
	return item
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Drain() []QueueItem {
	items := q.items
	q.items = nil
	return items
}
