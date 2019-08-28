// For more tutorials: https://bp.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package parse

import (
	"bufio"
	"io"

	"github.com/inancgumus/learngo/logparser/v5/pipe"
)

// Text parses text based log lines.
type Text struct {
	reader io.Reader
}

// FromText creates a text parser.
func FromText(r io.Reader) *Text {
	return &Text{reader: r}
}

// Each yields records from a text log.
func (p *Text) Each(yield func(pipe.Record)) error {
	defer readClose(p.reader)

	in := bufio.NewScanner(p.reader)

	for in.Scan() {
		r := new(record)

		if err := r.UnmarshalText(in.Bytes()); err != nil {
			return err
		}

		yield(r)
	}

	return in.Err()
}
