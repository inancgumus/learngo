// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package calc

import (
	"errors"
	"strconv"
)

// Parse ...
func Parse(snum string) (n float64, err error) {
	n, err = strconv.ParseFloat(snum, 64)
	if err != nil {
		// Don't loose the actual error for debugging
		err = errors.New("Please provide a valid number: " +
			err.Error())
	}
	return
}

// Do ...
func Do(a, b float64, op string) (res float64, err error) {
	switch op {
	case "+", "plus":
		op, res = "+", a+b
	case "-", "minus":
		op, res = "-", a-b
	case "*", "times":
		op, res = "*", a*b
	case "/", "div":
		op, res = "/", a/b
	case "%", "mod":
		res = float64(int(a) % int(b))
	default:
		return 0, errors.New("Wrong operation: '" + op + "'")
	}
	return
}
