package main

import (
	"fmt"
	"strconv"
)

// field helps for field parsing
type field struct{ err error }

// uatoi parses an unsigned integer string and saves the error.
// it assumes that the val is unsigned.
// for ease of usability: it returns an int instead of uint.
func (f *field) uatoi(name, val string) int {
	n, err := strconv.Atoi(val)
	if n < 0 || err != nil {
		f.err = fmt.Errorf("incorrect field -> %q = %q", name, val)
	}
	return n
}
