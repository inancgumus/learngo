// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strconv"
	"strings"
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
func parse(p *parser, line string) (r result) {
	// if there was an error do not continue
	if p.lerr != nil {
		return
	}

	// chain the parser's error to the result's
	r = parseLine(p, line)
	if p.lines++; p.lerr != nil {
		p.lerr = fmt.Errorf("line #%d: %s", p.lines, p.lerr)
	}
	return
}

// parse parses a single log line
func parseLine(p *parser, line string) (r result) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		// p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		p.lerr = fmt.Errorf("missing fields: %v", fields)
		return
	}

	r.domain = fields[0]
	r.visits, p.lerr = strconv.Atoi(fields[1])

	if r.visits < 0 || p.lerr != nil {
		p.lerr = fmt.Errorf("incorrect visits: %q", fields[1])
	}
	return
}

// summarizeParse summarizes the parsing results.
// Only use it after the parsing is done.
func summarizeParse(p *parser) *summary {
	return p.summary
}

// errParse returns the last error encountered
func errParse(p *parser) error {
	return p.lerr
}
