// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func textReader(r io.Reader) inputFunc {
	return func() ([]result, error) {
		// first: count the lines, so the parseText can create
		// enough buffer.
		var buf bytes.Buffer
		lines, err := countLines(io.TeeReader(r, &buf))
		if err != nil {
			return nil, err
		}

		return parseText(bufio.NewScanner(&buf), lines)
	}
}

// TODO: custom error type for line information
func parseText(in *bufio.Scanner, nlines int) ([]result, error) {
	res := make([]result, 0, nlines)

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

func countLines(r io.Reader) (int, error) {
	var (
		lines int
		buf   = make([]byte, 1024<<4) // read via 16 KB blocks
	)

	for {
		n, err := r.Read(buf)
		lines += bytes.Count(buf[:n], []byte{'\n'})

		if err == io.EOF {
			return lines, nil
		}

		if err != nil {
			return lines, err
		}
	}
}
