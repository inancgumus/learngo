// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

// parser parses the log file and generates a summary report.
type parser struct {
	summary *summary // summarizes the parsing results
	lines   int      // number of parsed lines (for the error messages)
	lerr    error    // the last error occurred
}

// new returns a new parsing state.
func newParser() *parser {
	return &parser{summary: newSummary()}
}

// parse parses a log line and adds it to the summary.
func (p *parser) parse(line string) {
	// if there was an error do not continue
	if p.lerr != nil {
		return
	}

	// chain the parser's error to the result's
	res, err := parseLine(line)
	if p.lines++; err != nil {
		p.lerr = fmt.Errorf("line #%d: %s", p.lines, err)
		return
	}

	p.summary.update(res)
}

// Summarize summarizes the parsing results.
// Only use it after the parsing is done.
func (p *parser) summarize() *summary {
	return p.summary
}

// Err returns the last error encountered
func (p *parser) err() error {
	return p.lerr
}
