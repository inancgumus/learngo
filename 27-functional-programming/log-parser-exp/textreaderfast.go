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
	"errors"
	"fmt"
	"io"
)

// this could be made faster.
// currently, it's 30-35% faster.
//
// so, what's different than the textreader?
//
// + creates the buffers specific to the input file/stdin size
// + manually parses the fields: instead of strings.Fields
// + gets the lines using scanner's Bytes() method: instead of Text()
// + uses a manual atoi
// +

func fastTextReader(r io.Reader) inputFunc {
	return func() ([]result, error) {
		// first: count the lines, so the parseText can create
		// enough buffer.
		var buf bytes.Buffer
		l, err := countLines(io.TeeReader(r, &buf))
		if err != nil {
			return nil, err
		}

		return fastParseText(bufio.NewScanner(&buf), l)
	}
}

func fastParseText(in *bufio.Scanner, nlines int) ([]result, error) {
	// needs to know the number of total lines in the file
	res := make([]result, 0, nlines)

	for l := 0; in.Scan(); l++ {
		r, err := fastParseFields(in.Bytes())

		if err != nil {
			return nil, fmt.Errorf("line %d: %v", l, err)
		}
		res = append(res, r)
	}

	return res, in.Err()
}

func fastParseFields(data []byte) (res result, err error) {
	var field int

	for i, last := 0, 0; i < len(data); i++ {
		done := len(data) == i+1

		if c := data[i]; c == ' ' || done {
			if done {
				i = len(data)
			}

			switch field {
			case 0:
				res.domain = string(data[last:i])
			case 1:
				res.page = string(data[last:i])
			case 2:
				res.visits, err = atoi(data[last:i])
			case 3:
				res.uniques, err = atoi(data[last:i])
			}

			if err != nil {
				return res, err
			}

			last = i + 1
			field++
		}
	}

	if field != 4 {
		return result{}, errors.New("wrong number of fields")
	}
	return res, nil
}

func atoi(input []byte) (int, error) {
	val := 0
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char < '0' || char > '9' {
			return 0, errors.New("invalid number")
		}
		val = val*10 + int(char) - '0'
	}
	return val, nil
}
