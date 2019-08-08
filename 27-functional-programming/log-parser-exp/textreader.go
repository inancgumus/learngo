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
	"io"
	"strings"
)

func textReader(r io.Reader) inputFn {
	return func() ([]result, error) {
		return parseText(bufio.NewScanner(r))
	}
}

// TODO: custom error type for line information
func parseText(in *bufio.Scanner) ([]result, error) {
	var res []result

	for l := 1; in.Scan(); l++ {
		fields := strings.Fields(in.Text())
		r, err := parseFields(fields)

		if err != nil {
			return nil, fmt.Errorf("line %d: %v", l, err)
		}
		res = append(res, r)
	}

	return res, in.Err()
}
