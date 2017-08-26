package mux

import (
	"log"
	"net/http"
	"time"
)

type WrapperFunc func(handlerFunc http.Handler) http.Handler

// Wraps handler func with slice of wrapper functions one by one.
func Wrap(h http.Handler, wrappers ...WrapperFunc) http.Handler {
	for _, w := range wrappers {
		h = w(h)
	}
	return h
}

// Logs requests
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

		h.ServeHTTP(w, r)
	})
}

// Creates wrapper functions that adds timeout to requests
func Timeout(timeout time.Duration) WrapperFunc {
	return func(h http.Handler) http.Handler {
		return http.TimeoutHandler(h, timeout, "timeout exceed")
	}
}
