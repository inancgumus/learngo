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
	"strconv"
	"strings"
)

const (
	validOps = "%*/+-"

	usageMsg       = "Usage: [op=" + validOps + "] [size]"
	sizeMissingMsg = "Size is missing"
	invalidOpMsg   = `Invalid operator.
Valid ops one of: ` + validOps

	invalidOp = -1
)

func main() {
	args := os.Args[1:]

	switch l := len(args); {
	case l == 1:
		fmt.Println(sizeMissingMsg)
		fallthrough
	case l < 1:
		fmt.Println(usageMsg)
		return
	}

	op := args[0]
	if strings.IndexAny(op, validOps) == invalidOp {
		fmt.Println(invalidOpMsg)
		return
	}

	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", op)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {
			var res int

			switch op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}

			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}
