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
	"time"
)

func main() {
	// time.Now() gets the current time
	// and in turn, .Hour() gets the current hour from it

	// h := time.Now().Hour()
	// fmt.Println("Current hour is", h)

	switch h := time.Now().Hour(); {
	case h >= 18: // 18 to 23
		fmt.Println("good evening")
	case h >= 12: // 12 to 18
		fmt.Println("good afternoon")
	case h >= 6: // 6 to 12
		fmt.Println("good morning")
	default: // 0 to 5
		fmt.Println("good night")
	}

	// h is not available here
	// it's bound to switch statement's block
	// fmt.Println(h)
}
