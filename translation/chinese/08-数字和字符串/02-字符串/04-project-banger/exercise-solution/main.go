// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"os"
	"strings"
)

// NOTE: You should always pass it at least one argument

func main() {
	msg := os.Args[1]

	// it's important to calculate things only once
	// so, do not call the repeat function twice
	// calling it once is enough
	marks := strings.Repeat("!", len(msg))
	s := marks + msg + marks
	s = strings.ToUpper(s)

	// you can also type this program more concisely
	// like this:
	//
	// msg := strings.ToUpper(os.Args[1])
	// marks := strings.Repeat("!", len(msg))
	// fmt.Println(marks + msg + marks)
}
