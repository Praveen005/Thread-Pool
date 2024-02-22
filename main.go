package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type pool struct {
	workQueue chan Job // A channel of jobs
	wg        sync.WaitGroup
}

func NewPool(workerCount int) *pool {
	pool := &pool{
		workQueue: make(chan Job),
	}

	pool.wg.Add(workerCount)

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

	for i := 0; i < 30; i++ {
		job := func() {
			time.Sleep(1 * time.Second)
			fmt.Println("job: completed")
		}
		pool.AddJob(job)
	}
	pool.Wait()
}
