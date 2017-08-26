package pool

import "fmt"

func getService(action Action) (ServiceFunc, error) {
	for prefix, resolver := range resolvers {
		if action.HasPrefix(prefix) {
			return resolver(action)
		}
	}

	return nil, fmt.Errorf("can not resolve service by action: %v", action)
}

// Creates job with given action and input and returns result.
func Dispatch(action Action, input interface{}) (interface{}, error) {
	worker := <-freeWorkers

	service, err := getService(action)
	if err != nil {
		return nil, fmt.Errorf("can not dispatch action: %v", err)
	}

	addJob(worker, input, service)

	jobResult := readResult(worker)

	return jobResult.result, jobResult.err
}
