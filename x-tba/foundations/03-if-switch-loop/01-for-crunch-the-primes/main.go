// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// go run . {1..100}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const usageMsg = `Type a couple of unique numbers.
Separate them with spaces.`

	// remember [1:] skips the first argument
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(usageMsg)
		return
	}

main:
	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			// skip non-numerics
			continue
		}

		switch {
		// prime
		case n == 2, n == 3:
			fmt.Print(n, " ")
			continue

		// not a prime
		case n <= 1, n%2 == 0, n%3 == 0:
			continue
		}

		for i, w := 5, 2; i*i <= n; {
			// not a prime
			if n%i == 0 {
				continue main
			}

			i += w
			w = 6 - w
		}

		// all checks ok: it's a prime
		fmt.Print(n, " ")
	}

	fmt.Println()
}
