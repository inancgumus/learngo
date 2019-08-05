package main

import (
	"fmt"
)

// result stores the parsed result for a domain
type result struct {
	domain  string
	page    string
	visits  int
	uniques int
}

// parseFields parses and returns the parsing result
func parseFields(fields []string) (r result, err error) {
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
