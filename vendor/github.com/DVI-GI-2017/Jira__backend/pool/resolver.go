package pool

import "log"

func getService(action Action) (service ServiceFunc) {
	for prefix, resolver := range resolvers {
		if action.HasPrefix(prefix) {
			return resolver(action)
		}
	}

	log.Panicf("can not resolve service by action: %v", action)
	return
}

// Creates job with given action and input and returns result.
func Dispatch(action Action, input interface{}) (interface{}, error) {
	jobs <- &job{
		input:   input,
		service: getService(action),
	}

	jobResult := <-results

	return jobResult.result, jobResult.err
}
