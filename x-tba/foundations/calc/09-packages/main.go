// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"

	"github.com/inancgumus/learngo/x-tba/foundations/calc/09-packages/calc"
)

func main() {
	// lesson: packaging

	// separate dependencies like getting and validating
	// user data as in here.
	//
	// we only put calc to another packages, not the other
	// stuff.
	//
	// so, we can reuse the same calc package and use it
	// over a web api if we want.

	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <number1> <operator> <number2>")
		return
	}

	var (
		a, b float64
		err  error
	)

	if a, err = calc.Parse(os.Args[1]); err != nil {
		fmt.Println(err)
		return
	}

	if b, err = calc.Parse(os.Args[3]); err != nil {
		fmt.Println(err)
		return
	}

	op := os.Args[2]
	res, err := calc.Do(a, b, op)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v %s %v = %v\n", a, op, b, res)
}
