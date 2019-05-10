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

// always put all the related things together as in here

// result stores metrics for a domain
// it uses the value mechanics,
// because it doesn't have to update anything
type result struct {
	domain    string
	visits    int
	timeSpent int
	// add more metrics if needed
}

// add adds the metrics of another result to itself and returns a new Result
func (r result) add(other result) result {
	return result{
		domain:    r.domain,
		visits:    r.visits + other.visits,
		timeSpent: r.timeSpent + other.timeSpent,
	}
}

// parse parses a single log line
func parseLine(line string) (r result, err error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return r, fmt.Errorf("missing fields: %v", fields)
	}

	f := new(field)
	r.domain = fields[0]
	r.visits = f.atoi("visits", fields[1])
	r.timeSpent = f.atoi("time spent", fields[2])
	return r, f.err
}

// field helps for field parsing
type field struct{ err error }

func (f *field) atoi(name, val string) int {
	n, err := strconv.Atoi(val)
	if n < 0 || err != nil {
		f.err = fmt.Errorf("incorrect %s: %q", name, val)
	}
	return n
}
