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
	sum     map[string]domain // visits per unique domain
	domains []string          // unique domain names
	total   int               // total visits to all domains
	lines   int               // number of parsed lines (for the error messages)
	lerr    error             // saves the last error occurred
}

// newParser creates and returns a new parser.
//
// bare func, it doesn't need to operate on a parser value.
func newParser() *parser {
	return &parser{sum: make(map[string]domain)}
}

// PUBLIC METHODS (API)

// add parses the given line and saves the result to the internal list of
// domains. it doesn't add the record when the parsing fails.
//
// WRITE METHOD
func (p *parser) add(line string) {
	// if there was a previous error do not add
	if p.lerr != nil {
		return
	}

	dom, err := p.parse(line)

	// store only the last error
	if err != nil {
		p.lerr = err
		return
	}

	p.push(dom)
}

// iterator returns two functions for iterating over domains.
// next = returns true when there are more domains to iterate on.
// cur  = returns the current domain
//
// READ METHOD
func (p *parser) iterator() (next func() bool, cur func() domain) {
	// remember the last received line
	var last int

	next = func() bool {
		defer func() { last++ }()
		return len(p.domains) > last
	}

	cur = func() domain {
		// return a copy so the caller cannot change it
		name := p.domains[last-1]
		return p.sum[name]
	}

	return
}

// error returns the last error occurred
//
// READ METHOD
func (p *parser) error() error {
	return p.lerr
}

// PRIVATE METHODS

// parse parses the given text and returns a domain struct
//
// WRITE METHOD
func (p *parser) parse(line string) (dom domain, err error) {
	p.lines++ // increase the parsed line counter (only write is here)

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

// push pushes the given domain to the internal list of domains.
// it also increases the total visits for all the domains.
//
// WRITE METHOD
func (p *parser) push(d domain) {
	name := d.name

	// collect the unique domains
	if _, ok := p.sum[name]; !ok {
		p.domains = append(p.domains, name)
	}

	p.total += d.visits
	d.visits += p.sum[name].visits
	p.sum[name] = d
}
