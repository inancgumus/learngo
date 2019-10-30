// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"errors"
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
	if err != nil || n < 0 {
		f.err = fmt.Errorf("incorrect field -> %q = %q", name, val)
	}
	return n
}

func atoi(input []byte) (int, error) {
	val := 0
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char < '0' || char > '9' {
			return 0, errors.New("invalid number")
		}
		val = val*10 + int(char) - '0'
	}
	return val, nil
}
