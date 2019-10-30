// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package report

import (
	"fmt"
	"strconv"
	"strings"
)

// always put all the related things together as in here

// Result stores metrics for a domain
// it uses the value mechanics,
// because it doesn't have to update anything
type Result struct {
	Domain    string `json:"domain"`
	Visits    int    `json:"visits"`
	TimeSpent int    `json:"time_spent"`
	// add more metrics if needed
}

// add adds the metrics of another Result to itself and returns a new Result
func (r Result) add(other Result) Result {
	return Result{
		Domain:    r.Domain,
		Visits:    r.Visits + other.Visits,
		TimeSpent: r.TimeSpent + other.TimeSpent,
	}
}

// parse parses a single log line
func parse(line string) (r Result, err error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return r, fmt.Errorf("missing fields: %v", fields)
	}

	f := new(field)
	r.Domain = fields[0]
	r.Visits = f.atoi("visits", fields[1])
	r.TimeSpent = f.atoi("time spent", fields[2])
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
