package main

import (
	"fmt"
	"time"
)

type PipelineStats struct {
	Total   int
	MinTook time.Duration
	MaxTook time.Duration
	AvgTook time.Duration
}

func computeStats(results []Result) PipelineStats {
	if len(results) == 0 {
		return PipelineStats{}
	}

	min := results[0].Took
	max := results[0].Took
	var sum time.Duration

	for _, r := range results {
		if r.Took < min {
			min = r.Took
		}
		if r.Took > max {
			max = r.Took
		}
		sum += r.Took
	}

	return PipelineStats{
		Total:   len(results),
		MinTook: min,
		MaxTook: max,
		AvgTook: sum / time.Duration(len(results)),
	}
}

func printStats(s PipelineStats) {
	fmt.Printf("jobs: %d | min: %s | max: %s | avg: %s\n", s.Total, s.MinTook, s.MaxTook, s.AvgTook)
}
