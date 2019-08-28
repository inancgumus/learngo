// For more tutorials: https://bp.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pipe

import (
	"bufio"
	"io"
)

// TextLog parses text based log lines.
type TextLog struct {
	reader io.Reader
}

// NewTextLog creates a text parser.
func NewTextLog(r io.Reader) *TextLog {
	return &TextLog{reader: r}
}

// Each yields records from a text log.
func (p *TextLog) Each(yield func(Record) error) error {
	defer readClose(p.reader)

	in := bufio.NewScanner(p.reader)

	for in.Scan() {
		r := new(Record)

		if err := r.UnmarshalText(in.Bytes()); err != nil {
			return err
		}

		if err := yield(*r); err != nil {
			return err
		}
	}

	return in.Err()
}
