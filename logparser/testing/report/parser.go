// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package report

import (
	"fmt"
)

// Parser parses the log file and generates a summary report.
type Parser struct {
	summary *Summary // summarizes the parsing results
	lines   int      // number of parsed lines (for the error messages)
	lerr    error    // the last error occurred
}

// New returns a new parsing state.
func New() *Parser {
	return &Parser{summary: newSummary()}
}

// Parse parses a log line and adds it to the summary.
func (p *Parser) Parse(line string) {
	// if there was an error do not continue
	if p.lerr != nil {
		return
	}

	// chain the parser's error to the result's
	res, err := parse(line)
	if p.lines++; err != nil {
		p.lerr = fmt.Errorf("line #%d: %s", p.lines, err)
		return
	}

	p.summary.update(res)
}

// Summarize summarizes the parsing results.
// Only use it after the parsing is done.
func (p *Parser) Summarize() *Summary {
	return p.summary
}

// Err returns the last error encountered
func (p *Parser) Err() error {
	return p.lerr
}
