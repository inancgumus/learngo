// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

/*
go run main.go < ./calculations.txt

If you're not on Linux or OS X, etc but on Windows,
Then, use Windows PowerShell to do the same thing.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// After functions, you'll see that how we're going to refactor this
// into a more readable version.

func main() {
	const (
		promptChar     = "> "
		errWrongOp     = "%s operation is not supported\n"
		errWrongFormat = "operation is not recognized\n"
		usage          = `
usage: number operation number
quit : type q to quit
examples:
  3 + 5
  5 - 3
`
	)

	fmt.Println(strings.TrimSpace(usage))

	for s := bufio.NewScanner(os.Stdin); ; {
		var (
			a, b, res float64
			op        string
		)

		fmt.Print(promptChar, " ")
		if !s.Scan() {
			break
		}

		_, err := fmt.Sscanf(s.Text(), "%f %s %f", &a, &op, &b)
		if err != nil {
			fmt.Fprintf(os.Stderr, errWrongFormat)
			continue
		}

		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		default:
			fmt.Printf(errWrongOp, op)
			continue
		}

		fmt.Printf("%g %s %g = %g\n", a, op, b, res)
	}

	fmt.Println("bye.")
}
