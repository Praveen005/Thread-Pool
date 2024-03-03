package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulating a job
type Job func()

// This is what the pool looks like, have a channel to store the jobs and a waitgroup associated with it
type pool struct {
	workQueue chan Job // A channel of jobs
	wg        sync.WaitGroup
}

// function to create a new pool
func NewPool(workerCount int) *pool {
	pool := &pool{
		workQueue: make(chan Job),
	}

	pool.wg.Add(workerCount)
	// Let's breakdown what's happening here:
	// It will spin up 'workerCount' number of goroutines, and each goroutine will try to iterate over the channel and finish the job,
	// Means, at one time concurrently 'workerCount' number of jobs would be getting executed by 'workerCount' number of goroutines, that's why you see the 'workerCount' number of print statements getting printed in the console in batches
	for i := 0; i < workerCount; i++ {
		go func() {
			defer pool.wg.Done()
			for job := range pool.workQueue {
				job()
			}
		}()
	}
	return pool
}

func (p *pool) AddJob(job Job) {
	p.workQueue <- job
}

func (p *pool) Wait() {
	close(p.workQueue)
	p.wg.Wait()
}

func main() {
	pool := NewPool(5)

	for i := 0; i < 20; i++ {
		// declaring j, to pass in a local copy, can see the diff. by removing it
		i := i
		job := func() {
			time.Sleep(2 * time.Second)
			fmt.Printf("job %d completed!\n", i)
		}
		pool.AddJob(job)
	}
	pool.Wait() // closes the channel and waits for all the goroutine to finish before exiting the main goroutine
}
