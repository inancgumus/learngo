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

func textReader(r io.Reader) ([]result, error) {
	in := bufio.NewScanner(r)
	return parseText(in)
}

func parseText(in *bufio.Scanner) ([]result, error) {
	var (
		results []result
		lines   int
	)

	for in.Scan() {
		lines++

		result, err := parseFields(strings.Fields(in.Text()))
		if err != nil {
			// TODO: custom error type for line information
			return nil, fmt.Errorf("line %d: %v", lines, err)
		}

		results = append(results, result)
	}

	if err := in.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
