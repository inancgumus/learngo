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
	"os"
	"strconv"
)

func main() {
	// lesson: many - split this

	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <number1> <operator> <number2>")
		return
	}

	var (
		a, b float64
		err  error
	)

	if a, err = parse(os.Args[1]); err != nil {
		fmt.Println(err)
		return
	}

	if b, err = parse(os.Args[3]); err != nil {
		fmt.Println(err)
		return
	}

	op := os.Args[2]
	res, err := calc(a, b, op)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v %s %v = %v\n", a, op, b, res)
}

func parse(snum string) (n float64, err error) {
	n, err = strconv.ParseFloat(snum, 64)
	if err != nil {
		err = errors.New("Please provide a valid number")
	}
	return
}

func calc(a, b float64, op string) (res float64, err error) {
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
