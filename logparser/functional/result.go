// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
)

const fieldsLength = 4

// result stores the parsed result for a domain
type result struct {
	domain  string
	page    string
	visits  int
	uniques int
}

// add adds the metrics of another result
func (r result) add(other result) result {
	r.visits += other.visits
	r.uniques += other.uniques
	return r
}

// parseFields parses and returns the parsing result
func parseFields(line string) (r result, err error) {
	fields := strings.Fields(line)

	if len(fields) != fieldsLength {
		return r, fmt.Errorf("wrong number of fields %q", fields)
	}

	r.domain = fields[0]
	r.page = fields[1]

	f := new(field)
	r.visits = f.uatoi("visits", fields[2])
	r.uniques = f.uatoi("uniques", fields[3])

	return r, f.err
}

func fastParseFields(data []byte) (res result, err error) {
	const separator = ' '

	var findex int

	for i, j := 0, 0; i < len(data); i++ {
		c := data[i]
		last := len(data) == i+1

		if c != separator && !last {
			continue
		}

		if last {
			i = len(data)
		}

		switch fval := data[j:i]; findex {
		case 0:
			res.domain = string(fval)
		case 1:
			res.page = string(fval)
		case 2:
			res.visits, err = atoi(fval)
		case 3:
			res.uniques, err = atoi(fval)
		}

		if err != nil {
			return res, err
		}

		j = i + 1
		findex++
	}

	if findex != fieldsLength {
		err = fmt.Errorf("wrong number of fields %q", data)
	}
	return res, err
}
