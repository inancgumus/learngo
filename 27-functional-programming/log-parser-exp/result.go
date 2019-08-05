package main

import (
	"fmt"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain  string
	page    string
	visits  int
	uniques int
}

// parseLine parses a log line and returns the parsed result with an error
func parseLine(line string) (r result, err error) {
	fields := strings.Fields(line)
	if len(fields) != 4 {
		return r, fmt.Errorf("wrong number of fields -> %v", fields)
	}

	r.domain = fields[0]
	r.page = fields[1]

	f := new(field)
	r.visits = f.uatoi("visits", fields[2])
	r.uniques = f.uatoi("uniques", fields[3])

	return r, f.err
}

// add adds the metrics of another result to the result
func (r result) add(other result) result {
	r.visits += other.visits
	r.uniques += other.uniques
	return r
}
