package pool

import (
	"runtime"

	"github.com/DVI-GI-2017/Jira__backend/db"
)

// Starts worker pool
func init() {
	numCPU := runtime.NumCPU()

	freeWorkers = make(chan int, numCPU)
	jobs = make(map[int](chan job), numCPU)
	results = make(map[int](chan jobResult), numCPU)

	for id := 0; id < numCPU; id++ {
		jobs[id] = make(chan job, 100)
		results[id] = make(chan jobResult, 100)

		go worker(id)
	}
}

// Helper maps of channels
var jobs map[int](chan job)
var results map[int](chan jobResult)

// Channel for free workers ids
var freeWorkers chan int

// Type for jobs
type job struct {
	service ServiceFunc
	input   interface{}
}

// Process job with self-contained input and given data source
func (j job) process() (result interface{}, err error) {
	source := db.Copy()
	defer source.Close()

	return j.service(source, j.input)
}

// Type for results
type jobResult struct {
	err    error
	result interface{}
}

// Reads from associated jobs channel and writes to associated results channel
func worker(id int) {
	freeWorkers <- id

	for job := range jobs[id] {
		result, err := job.process()
		results[id] <- jobResult{result: result, err: err}

		// Add to free workers
		freeWorkers <- id
	}
}

// Adds job to channel
func addJob(id int, input interface{}, service ServiceFunc) {
	jobs[id] <- job{input: input, service: service}
}

// Read result from channel
func readResult(id int) jobResult {
	return <-results[id]
}
