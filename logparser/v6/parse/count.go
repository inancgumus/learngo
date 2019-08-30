// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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

// Last counted record number.
func (c *Count) Last() int {
	return c.count - 1
}

// Parse increments the counter.
func (c *Count) Parse() bool {
	c.count++
	return c.Parser.Parse()
}

// Err returns the first error that was encountered by the Log.
func (c *Count) Err() (err error) {
	if err = c.Parser.Err(); err != nil {
		err = fmt.Errorf("record #%d: %v", c.Last()+1, err)
	}
	return
}

// You don't need to implement the Value() method.
// Thanks to interface embedding.
