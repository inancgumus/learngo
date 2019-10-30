// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package parse

import "fmt"

// Count the parsed records.
type Count struct {
	// Parser is wrapped by Count to count the parsed records.
	Parser
	count int
}

// CountRecords creates a record counter that wraps a parser.
func CountRecords(p Parser) *Count {
	return &Count{Parser: p}
}

// Parse increments the counter.
func (c *Count) Parse() bool {
	c.count++
	return c.Parser.Parse()
}

// Err returns the first error that was encountered by the Log.
func (c *Count) Err() (err error) {
	err = c.Parser.Err()
	if err != nil {
		err = fmt.Errorf("record #%d: %v", c.count, err)
	}
	return
}

// You don't need to implement the Value() method.
// Thanks to interface embedding.
