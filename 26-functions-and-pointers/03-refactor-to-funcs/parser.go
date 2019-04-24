package main

import (
	"fmt"
	"strconv"
	"strings"
)

// domain represents a domain log record
type domain struct {
	name   string
	visits int
}

// parser parses a log file and provides an iterator to iterate upon the domains
//
// the parser struct is carefully crafted to be usable using its zero values except the map field
type parser struct {
	sum     map[string]int // visits per unique domain
	domains []domain       // unique domain names
	total   int            // total visits to all domains
	lines   int            // number of parsed lines (for the error messages)
}

// newParser creates and returns a new parser.
func newParser() parser {
	return parser{sum: make(map[string]int)}
}

// parse parses the given text and returns a domain struct
func parse(p parser, line string) (dom domain, err error) {
	// var dom domain
	// var err error

	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}
	name, visits := fields[0], fields[1]

	n, err := strconv.Atoi(visits)
	if n < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", visits, p.lines)
		return
	}

	return domain{name: name, visits: n}, nil
}
