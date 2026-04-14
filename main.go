package main

import (
	"fmt"
	"time"
)

func main() {
	task := NewTask(
		"Implement prism core",
		"Build the main refraction logic for go-prism",
		PriorityHigh,
		time.Now().Add(48*time.Hour),
	)

	fmt.Println(task)

	task.Complete()
	fmt.Println("After completion:", task)
}
