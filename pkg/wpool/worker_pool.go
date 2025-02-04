package wpool

import (
	"context"
	"fmt"
	"sync"
)

type WorkerPool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
	wg           sync.WaitGroup
	Done         chan struct{}
}

func New(wcount int) *WorkerPool {
	return &WorkerPool{
		workersCount: wcount,
		jobs:         make(chan Job, wcount*2),
		results:      make(chan Result, wcount),
		Done:         make(chan struct{}),
		wg:           sync.WaitGroup{},
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result, id int) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			fmt.Printf("[Worker %d] Processing job ID: %s\n", id, job.Descriptor.ID)

			results <- job.execute(ctx)
		case <-ctx.Done():
			fmt.Printf("[Worker %d] Cancelled. Detail: %v\n", id, ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 1; i <= wp.workersCount; i++ {
		wp.wg.Add(1)
		go worker(ctx, &wp.wg, wp.jobs, wp.results, i)
	}

	wp.wg.Wait()
	close(wp.Done)
	close(wp.results)
}

func (wp *WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp *WorkerPool) GenerateJobs(jobs []Job) {
	for _, job := range jobs {
		wp.jobs <- job
	}

	close(wp.jobs)
}
