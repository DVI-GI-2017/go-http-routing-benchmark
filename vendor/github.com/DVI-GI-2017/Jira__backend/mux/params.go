package mux

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

// Type for http post body
type PostBody []byte // Byte array with request body

// Get params stands for "query params"
type GetParams map[string]string

// Extracts "params" from request
func Params(req *http.Request) *params {
	return req.Context().Value("params").(*params)
}

// Puts params in request context and returns new request
func putParams(req *http.Request, params *params) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), "params", params))
}

type params struct {
	Query      GetParams
	Body       PostBody
	PathParams PathParams
}

// Creates new params
func newParams(request *http.Request, pattern *regexp.Regexp, path string) (*params, error) {
	var body []byte
	if request.Body != nil {
		newBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		body = newBody
	}

	return &params{
		Query:      valuesToGetParams(request.URL.Query()),
		Body:       body,
		PathParams: extractPathParams(pattern, path),
	}, nil
}

// Converts url.Url.Query() from "Values" (map[string][]string)
// to "getParams" (map[string]string)
func valuesToGetParams(values url.Values) GetParams {
	params := make(map[string]string)
	for key := range values {
		params[key] = values.Get(key)
	}
	return params
}

// Example: url "/api/v1/users/599a49bacdf43b817eeea57b" and pattern `/api/v1/users/:id`
// path params = {"id": "599a49bacdf43b817eeea57b"}
type PathParams map[string]string

// Extract path params from path
func extractPathParams(pattern *regexp.Regexp, path string) PathParams {
	match := pattern.FindStringSubmatch(path)
	result := make(PathParams)

	for i, name := range pattern.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	return result
}
