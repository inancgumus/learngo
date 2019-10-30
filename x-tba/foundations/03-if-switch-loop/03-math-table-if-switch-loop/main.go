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
	validOps       = "* / + - mul div add sub"
	usageMsg       = "Usage: [valid ops: " + validOps + "] [size]"
	sizeMissingMsg = "Size is missing\n" + usageMsg
	invalidOpMsg   = `Invalid operator.
Valid ops one of: ` + validOps
)

func main() {
	// CHECK THE ARGUMENTS
	args := os.Args[1:]
	if l := len(args); l == 1 {
		fmt.Println(sizeMissingMsg)
		return
	} else if l < 1 {
		fmt.Println(usageMsg)
		return
	}

	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}

	// CHECK THE VALIDITY OF THE OP
	op, ops := args[0], strings.Fields(validOps)

	var ok bool
	for _, o := range ops {
		if strings.ToLower(o) == op {
			ok = true
			break
		}
	}

	if !ok {
		fmt.Println(invalidOpMsg)
		return
	}

	// PRINT THE TABLE

	// HEADER
	fmt.Printf("%5s", op)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// CELLS
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {
			var res int

			switch op {
			default:
				fallthrough // default is multiplication
			case "*", "mul":
				res = i * j
			case "+", "add":
				res = i + j
			case "-", "sub":
				res = i - j
			case "%", "mod":
				if j == 0 {
					// continue // will continue the loop
					break // breaks the switch
				}
				res = i % j
			}

			// break // breaks the loop

			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}
