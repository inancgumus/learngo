// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package parse

import (
	"encoding/json"
	"io"

	"github.com/inancgumus/learngo/logparser/v6/logly/record"
)

// JSONParser parses json records.
type JSONParser struct {
	in   *json.Decoder
	err  error          // last error
	last *record.Record // last parsed record
}

// JSON creates a json parser.
func JSON(r io.Reader) *JSONParser {
	return &JSONParser{
		in:   json.NewDecoder(r),
		last: new(record.Record),
	}
}

// Parse the next line.
func (p *JSONParser) Parse() bool {
	if p.err != nil {
		return false
	}

	p.last.Reset()

	err := p.in.Decode(&p.last)
	if err == io.EOF {
		return false
	}

	p.err = err

	return err == nil
}

// Value returns the most recent record parsed by a call to Parse.
func (p *JSONParser) Value() record.Record {
	return *p.last
}

// Err returns the first error that was encountered by the Log.
func (p *JSONParser) Err() error {
	return p.err
}
