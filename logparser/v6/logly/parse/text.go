// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package parse

import (
	"bufio"
	"io"

	"github.com/inancgumus/learngo/logparser/v6/logly/record"
)

// TextParser parses text based log lines.
type TextParser struct {
	in   *bufio.Scanner
	err  error          // last error
	last *record.Record // last parsed record
}

// Text creates a text parser.
func Text(r io.Reader) *TextParser {
	return &TextParser{
		in:   bufio.NewScanner(r),
		last: new(record.Record),
	}
}

// Parse the next line.
func (p *TextParser) Parse() bool {
	if p.err != nil {
		return false
	}
	if !p.in.Scan() {
		return false
	}

	p.err = p.last.FromText(p.in.Bytes())

	return true
}

// Value returns the most recent record parsed by a call to Parse.
func (p *TextParser) Value() record.Record {
	return *p.last
}

// Err returns the first error that was encountered by the Log.
func (p *TextParser) Err() error {
	return p.err
}
