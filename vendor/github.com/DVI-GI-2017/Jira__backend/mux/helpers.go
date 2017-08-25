package mux

import (
	"errors"
	"fmt"
	"regexp"
)

var paramRegexp = regexp.MustCompile(`:([[:lower:]]|_)+`)

// Converts patterns like "/users/:id" to "/users/(?P<id>\d+)"
func convertSimplePatternToRegexp(pattern string) string {
	patternWithParams := paramRegexp.ReplaceAllStringFunc(pattern, func(param string) string {
		return fmt.Sprintf(`(?P<%s>[[:xdigit:]]{24})`, param[1:])
	})

	return fmt.Sprintf(`^%s/?$`, patternWithParams)
}

// Return path relative to "base"
func relativePath(base string, absolute string) (string, error) {
	baseLen := len(base)
	absoluteLen := len(absolute)

	if absoluteLen < baseLen {
		return "", errors.New("absolute len shorter than base len")
	}

	if absolute[:baseLen] != base {
		return "", errors.New("absolute path doesn't start with base path")
	}

	return absolute[baseLen:], nil
}
