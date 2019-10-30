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
)

func main() {
	// remember [1:] skips the first argument

main:
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil {
			// skip non-numerics
			continue
		}

		switch {
		// prime
		case n == 2 || n == 3:
			fmt.Print(n, " ")
			continue

		// not a prime
		case n <= 1 || n%2 == 0 || n%3 == 0:
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
