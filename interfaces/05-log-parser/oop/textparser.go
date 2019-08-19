// For more tutorials: https://bp.learngoprogramming.com
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
	"strconv"
	"strings"
)

type textParser struct {
	r io.Reader
}

func newTextParser(r io.Reader) *textParser {
	return &textParser{r: r}
}

func (p *textParser) parse(handle resultFn) error {
	defer readClose(p.r)

	var (
		l  = 1
		in = bufio.NewScanner(p.r)
	)

	for in.Scan() {
		r, err := parseFields(in.Text())
		if err != nil {
			return fmt.Errorf("line %d: %v", l, err)
		}

		handle(r)
		l++
	}

	return in.Err()
}

func parseFields(s string) (r result, err error) {
	fields := strings.Fields(s)
	if len(fields) != fieldsLength {
		return r, fmt.Errorf("wrong number of fields %q", fields)
	}

	r.domain, r.page = fields[0], fields[1]

	r.visits, err = strconv.Atoi(fields[2])
	if err != nil || r.visits < 0 {
		return r, fmt.Errorf("wrong input %q", fields[2])
	}

	r.uniques, err = strconv.Atoi(fields[3])
	if err != nil || r.uniques < 0 {
		return r, fmt.Errorf("wrong input %q", fields[3])
	}

	return r, nil
}
