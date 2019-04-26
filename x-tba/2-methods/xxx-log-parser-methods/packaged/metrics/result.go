// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package metrics

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
	Domain string
	Visits int
	// add more metrics if needed
}

// add adds the metrics of another Result to itself and returns a new Result
func (r result) add(other result) result {
	return result{
		Domain: r.Domain,
		Visits: r.Visits + other.Visits,
	}
}

// parse parses a single log line
func parse(line string) (parsed result, err error) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v", fields)
		return
	}

	parsed.Domain = fields[0]

	parsed.Visits, err = strconv.Atoi(fields[1])
	if parsed.Visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q", fields[1])
	}

	return
}
