// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "io"

// counter counts the total number of bytes read.
type counter struct {
	r io.Reader
	n int
}

// newCounter wraps `r` and returns a *counter.
func newCounter(r io.Reader) *counter {
	return &counter{r: r}
}

// Total bytes read.
func (c *counter) total() int {
	return c.n
}

// Read counts the number of bytes read from the underlying reader.
func (c *counter) Read(p []byte) (n int, err error) {
	n, err = c.r.Read(p) // read from the underlying reader
	if n > 0 {
		c.n += n // sum the number of bytes read with the total bytes counted
	}
	return n, err // return the number of bytes read and an error
}
