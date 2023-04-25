package main

import (
	"log"

	"github.com/kirill-scherba/jobs"
)

func main() {

	const (
		numWorkers = 5
		numJobs    = 10
	)

	// Set Log in microseconds
	log.SetFlags(log.Default().Flags() | log.Lmicroseconds)

	// Create jobs worker pool with Number of Workers equal to numWorkers and
	// Jobs Queue equal to numWorkers
	j := jobs.New(numWorkers, numWorkers, func(in interface{}) (out interface{}) {
		out = in
		return
	})

	// Add jobs to worker pool
	go func() {
		for job := 1; job <= numJobs; job++ {
			j.Add(job)
		}
		j.AddDone()
	}()

	// Collect jobs result
	for res := range j.Results() {
		log.Printf("res: %v", res)
	}
}
