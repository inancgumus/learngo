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

type processFn func(r result)

func scan(process processFn) error {
	var (
		l  = 1
		in = bufio.NewScanner(os.Stdin)
	)

	for in.Scan() {
		r, err := parseResult(in.Text())
		if err != nil {
			return fmt.Errorf("line %d: %v", l, err)
		}

		l++

		process(r)
	}
	return in.Err()
}
