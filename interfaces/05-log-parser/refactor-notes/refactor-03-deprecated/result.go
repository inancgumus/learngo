// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const fieldsLength = 4

// result stores the parsed result for a domain
type result struct {
	domain, page    string
	visits, uniques int
	// add more metrics if needed
}

// parseResult from a log line
func parseResult(line string) (r result, err error) {
	fields := strings.Fields(line)
	if len(fields) != fieldsLength {
		return r, fmt.Errorf("wrong input: %v", fields)
	}

	r.domain = fields[0]
	r.page = fields[1]

	r.visits, err = strconv.Atoi(fields[2])
	if err != nil || r.visits < 0 {
		return r, fmt.Errorf("wrong input: %q", fields[2])
	}

	r.uniques, err = strconv.Atoi(fields[3])
	if err != nil || r.uniques < 0 {
		return r, fmt.Errorf("wrong input: %q", fields[3])
	}

	return r, nil
}

// addResult to another one
func addResult(r, other result) result {
	r.visits += other.visits
	r.uniques += other.uniques
	return r
}
