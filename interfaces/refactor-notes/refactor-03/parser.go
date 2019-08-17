// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"fmt"
	"os"
)

// pipeline determines the behavior of log processing
type pipeline struct {
}

// pipe the log lines through funcs and produce a result
func pipe(opts pipeline) ([]result, error) {
	var (
		l   = 1
		in  = bufio.NewScanner(os.Stdin)
		sum = make(map[string]result)
	)

	// parse the log lines
	for in.Scan() {
		r, err := parseResult(in.Text())
		if err != nil {
			return nil, fmt.Errorf("line %d: %v", l, err)
		}

		l++

		// group the log lines by domain
		sum[r.domain] = addResult(r, sum[r.domain])
	}

	// collect the grouped results
	res := make([]result, 0, len(sum))
	for _, r := range sum {
		res = append(res, r)
	}

	return res, in.Err()
}
