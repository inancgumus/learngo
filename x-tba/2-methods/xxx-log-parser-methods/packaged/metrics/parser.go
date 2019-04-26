// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package metrics

import (
	"fmt"
)

// Parser keep tracks of the parsing
type Parser struct {
	lines int   // number of parsed lines (for the error messages)
	lerr  error // the last error occurred
}

// Parsed wraps a result for generating parser error
type Parsed struct {
	result       // use struct embedding
	err    error // inject an error
}

// NewParser returns a new parser
func NewParser() *Parser {
	return new(Parser)
}

// Parse parses a log line and returns a result with an injected error
func (l *Parser) Parse(line string) (parsed Parsed) {
	// always set the error
	defer func() { parsed.err = l.lerr }()

	// if there was an error do not continue
	if l.lerr != nil {
		return
	}

	// chain the parser's error to the result's
	res, err := parse(line)
	if l.lines++; err != nil {
		l.lerr = fmt.Errorf("%s: (line #%d)", err, l.lines)
	}
	return Parsed{result: res}
}

// Err returns the last error encountered
func (l *Parser) Err() error {
	return l.lerr
}
