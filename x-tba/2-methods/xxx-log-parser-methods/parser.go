package main

import (
	"fmt"
)

// parser keep tracks of the parsing
type parser struct {
	lines int   // number of parsed lines (for the error messages)
	lerr  error // the last error occurred
}

// parserResult wraps a result for generating parser error
type parserResult struct {
	result       // use struct embedding
	err    error // inject an error
}

// parse parses a log line and returns a result with an injected error
func (p *parser) parse(line string) (parsed parserResult) {
	// always set the error
	defer func() { parsed.err = p.lerr }()

	// if there was an error do not continue
	if p.lerr != nil {
		return
	}

	// chain the parser's error to the result's
	res, err := parseLine(line)
	if p.lines++; err != nil {
		p.lerr = fmt.Errorf("%s: (line #%d)", err, p.lines)
	}
	return parserResult{result: res}
}
