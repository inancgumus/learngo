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
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	u, p := args[1], args[2]

	if u != "jack" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if p != "1888" {
		fmt.Printf("Invalid password for %q.\n", u)
	} else {
		fmt.Printf("Access granted to %q.\n", u)
	}
}
