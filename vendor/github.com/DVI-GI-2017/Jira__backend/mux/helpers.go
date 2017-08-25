package mux

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Supported types
const stringType = `(?P<%s>[[:alnum:]]+)`
const hexType = `(?P<%s>[[:xdigit:]]{24})`
const intType = `(?P<%s>[[:digit:]]+)`

var paramRegexp = regexp.MustCompile(`(str|hex|int)??:((?:[[:lower:]]|_)+)`)

// Converts patterns like "/users/id:hex" to real regexps
func convertSimplePatternToRegexp(pattern string) string {
	patternWithParams := paramRegexp.ReplaceAllStringFunc(pattern, func(param string) string {
		paramParts := strings.Split(param, ":")

		if len(paramParts) == 1 || len(paramParts[0]) == 0 {
			return fmt.Sprintf(stringType, paramParts[1])
		}
		if len(paramParts) == 2 {
			fmtString, err := getPatternByType(paramParts[0])

			if err != nil {
				log.Panicf("wrong pattern format %s: %v", param, err)
				return ""
			}

			return fmt.Sprintf(fmtString, paramParts[1])
		}

		log.Panicf("wrong pattern format %s", param)
		return ""
	})

	return fmt.Sprintf(`^%s/?$`, patternWithParams)
}

func getPatternByType(name string) (string, error) {
	switch name {
	case "hex":
		return hexType, nil
	case "int":
		return intType, nil
	case "str":
		return stringType, nil
	default:
		return "", fmt.Errorf("can not find type with name '%s'", name)
	}
}

// Return path relative to "base"
func relativePath(base string, absolute string) (string, error) {
	if len(base) == 0 {
		return absolute, nil
	}

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
