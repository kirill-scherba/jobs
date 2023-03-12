package jobs

import "testing"

func TestJobs(t *testing.T) {

	const (
		numWorkers = 5
		numJobs    = 10
	)

	// Create jobs worker pool with Number of Workers equal to numWorkers and
	// Jobs Queue equal to numWorkers
	j := New(numWorkers, numWorkers, func(in interface{}) (out interface{}) {
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
		t.Logf("res: %v", res)
	}
}
