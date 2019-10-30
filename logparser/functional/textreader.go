// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"io"
)

func textReader(r io.Reader) inputFn {
	return func(process processFn) error {
		var (
			l  = 1
			in = bufio.NewScanner(r)
		)

		for in.Scan() {
			r, err := fastParseFields(in.Bytes())
			// r, err := parseFields(in.Text())
			if err != nil {
				return fmt.Errorf("line %d: %v", l, err)
			}

			process(r)
			l++
		}

		if c, ok := r.(io.Closer); ok {
			c.Close()
		}
		return in.Err()
	}
}
