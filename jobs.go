// Copyright 2023 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Jobs create and use a worker pool to limit the number of Goroutines that run
// at the same time in Go application.
package jobs

import "sync"

// Jobs method receiver
type Jobs struct {
	jobs    chan interface{} // jobs input cannel
	results chan interface{}
	process JobsFunc
	wg      *sync.WaitGroup
}

// JobsFunc is user defined function which executes for each job
type JobsFunc func(in interface{}) interface{}

// New creates a new Jobs object
func New(numWorkers int, jobsQueueSize int, process JobsFunc) (j *Jobs) {

	j = new(Jobs)

	j.jobs = make(chan interface{}, jobsQueueSize)
	j.results = make(chan interface{}, jobsQueueSize)
	j.process = process
	j.wg = new(sync.WaitGroup)

	// Create workers
	for w := 1; w <= numWorkers; w++ {
		j.wg.Add(1)
		go j.worker(w)
	}

	// Wait all jobs done
	go func() {
		j.wg.Wait()
		close(j.results)
	}()

	return
}

// Worker starts worker goroutine
func (j Jobs) worker(id int) {
	for job := range j.jobs {
		j.results <- j.process(job)
	}
	j.wg.Done()
}

// Add jobs to worker pool
func (j Jobs) Add(jobs ...interface{}) {
	for _, job := range jobs {
		j.jobs <- job
	}
}

// Indicate that all jobs was added
func (j Jobs) AddDone() {
	close(j.jobs)
}

// Results get jobs results channel
func (j Jobs) Results() chan interface{} {
	return j.results
}
