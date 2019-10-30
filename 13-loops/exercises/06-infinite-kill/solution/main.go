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
	"math/rand"
	"time"
)

func main() {
	for {
		var c string

		switch rand.Intn(4) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(time.Millisecond * 250)
	}
}
