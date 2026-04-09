package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	q := NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Printf("Queue size: %d\n", q.Size())

	for _, item := range q.Drain() {
		fmt.Printf("  [%s] value=%d\n", item.ID, item.Value)
	}

	fmt.Printf("Queue size after drain: %d\n", q.Size())
}
