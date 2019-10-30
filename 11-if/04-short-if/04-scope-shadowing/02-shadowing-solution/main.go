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
	// n and err belongs to the main function now
	// not only to the if statement below
	var (
		n   int
		err error
	)

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		// here else if uses the main func's n and err
		// variables instead of its own

		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		// assigns the result back into main func's
		// n variable
		n *= 2

		fmt.Printf("%s * 2 is %d\n", a[1], n)
	}

	// if statement has calculated the result using
	// the main func's n variable
	//
	// so, that's why it prints the correct result here
	fmt.Println("n is", n)
}
