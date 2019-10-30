// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	var pattern string
	if args := os.Args[1:]; len(args) == 1 {
		pattern = args[0]
	}

	for in.Scan() {
		s := in.Text()
		if strings.Contains(s, pattern) {
			fmt.Println(s)
		}
	}
}
