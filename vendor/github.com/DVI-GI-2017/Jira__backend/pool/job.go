package pool

import "github.com/DVI-GI-2017/Jira__backend/db"

var jobs = make(chan *job, 512)

type job struct {
	service ServiceFunc
	input   interface{}
}

// Process job with self-contained input and given data source
func (j job) process() {
	source := db.Copy()
	defer source.Close()

	result, err := j.service(source, j.input)

	results <- &jobResult{
		err:    err,
		result: result,
	}
}
