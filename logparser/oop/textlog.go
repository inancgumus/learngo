// For more tutorials: https://bp.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"io"
)

type textLog struct {
	reader io.Reader
}

func newTextLog(r io.Reader) *textLog {
	return &textLog{reader: r}
}

func (p *textLog) each(yield recordFn) error {
	defer readClose(p.reader)

	in := bufio.NewScanner(p.reader)

	for in.Scan() {
		r := new(record)

		if err := r.UnmarshalText(in.Bytes()); err != nil {
			return err
		}

		yield(*r)
	}

	return in.Err()
}
