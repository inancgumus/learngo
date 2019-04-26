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
	domain string
	visits int
	// add more metrics if needed
}

// add adds the metrics of another result to itself and returns a new result
func (r result) add(other result) result {
	return result{
		domain: r.domain,
		visits: r.visits + other.visits,
	}
}

// parseLine parses a single result line
func parseLine(line string) (p result, err error) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v", fields)
		return
	}

	p.domain = fields[0]

	p.visits, err = strconv.Atoi(fields[1])
	if p.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q", fields[1])
	}

	return
}
