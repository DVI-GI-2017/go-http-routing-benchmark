package pool

import (
	"runtime"
)

type jobResult struct {
	err    error
	result interface{}
}

var results = make(chan *jobResult, 512)

func InitWorkers() {
	for id := 0; id < runtime.NumCPU(); id++ {
		go worker()
	}
}

func worker() {
	for job := range jobs {
		job.process()
	}
}
