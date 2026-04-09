package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("First Commit")
	fmt.Println("Second Commit")
	printColored("info", "colored output via fatih/color")

	jobs := []Job{
		{ID: 1, Value: 3},
		{ID: 2, Value: 7},
		{ID: 3, Value: 5},
		{ID: 4, Value: 2},
	}

	results, err := runPipeline(context.Background(), jobs, 2)
	if err != nil {
		fmt.Println("pipeline error:", err)
		return
	}

	for _, r := range results {
		fmt.Printf("job %d: %d² = %d (took %s)\n", r.JobID, jobs[r.JobID-1].Value, r.Output, r.Took)
	}

	printStats(computeStats(results))
}
