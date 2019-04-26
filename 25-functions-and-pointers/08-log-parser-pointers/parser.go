package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: add add() func
// TODO: add error handling (variadics)
// TODO: add iterator func values
// TODO: add summarizer to main()

// domain represents a single domain log record
type domain struct {
	name   string
	visits int
}

// parser parses a log file and provides an iterator to iterate upon the domains
//
// the parser struct is carefully crafted to be usable using its zero values except the map field
type parser struct {
	// sum     map[string]int // visits per unique domain
	// domains []domain       // unique domain names
	sum     map[string]domain // visits per unique domain
	domains []string          // unique domain names

	total int   // total visits to all domains
	lines int   // number of parsed lines (for the error messages)
	lerr  error // saves the last error occurred
}

// newParser creates and returns a new parser.
func newParser() *parser {
	// return &parser{sum: make(map[string]int)}
	return &parser{sum: make(map[string]domain)}
}

// add parses the given line and saves the result to the internal list of
// domains. it doesn't add the record when the parsing fails.
func add(p *parser, line string) {
	// if there was a previous error do not add
	if p.lerr != nil {
		return
	}

	dom, err := parse(p, line)

	// store only the last error
	if err != nil {
		p.lerr = err
		return
	}

	push(p, dom)
}

// parse parses the given text and returns a domain struct
func parse(p *parser, line string) (dom domain, err error) {
	p.lines++

	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	dom.name = fields[0]

	dom.visits, err = strconv.Atoi(fields[1])
	if dom.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}

	return
}

// push pushes the given domain to the internal list of domains.
// it also increases the total visits for all the domains.
func push(p *parser, d domain) {
	// TODO:
	// if _, ok := p.sum[d.name]; !ok {
	// 	p.domains = append(p.domains, d)
	// }

	// p.sum[d.name] += d.visits
	// p.total += d.visits
	name := d.name

	// collect the unique domains
	if _, ok := p.sum[name]; !ok {
		p.domains = append(p.domains, name)
	}

	p.total += d.visits
	d.visits += p.sum[name].visits
	p.sum[name] = d
}
