package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Output int
	Took   time.Duration
}

func process(job Job) (Result, error) {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	if job.Value < 0 {
		return Result{}, fmt.Errorf("job %d: negative value not allowed", job.ID)
	}

	return Result{
		JobID:  job.ID,
		Output: job.Value * job.Value,
		Took:   time.Since(start),
	}, nil
}

func runPipeline(ctx context.Context, jobs []Job, workers int) ([]Result, error) {
	jobCh := make(chan Job, len(jobs))
	resultCh := make(chan Result, len(jobs))

	for _, j := range jobs {
		jobCh <- j
	}
	close(jobCh)

	g, ctx := errgroup.WithContext(ctx)

	for range workers {
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case job, ok := <-jobCh:
					if !ok {
						return nil
					}
					res, err := process(job)
					if err != nil {
						return err
					}
					resultCh <- res
				}
			}
		})
	}

	go func() {
		g.Wait()
		close(resultCh)
	}()

	var results []Result
	for r := range resultCh {
		results = append(results, r)
	}

	return results, g.Wait()
}
