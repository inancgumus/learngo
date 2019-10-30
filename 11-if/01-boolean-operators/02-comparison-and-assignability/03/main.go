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
	"strings"
)

func main() {
	speed := 100 // #1
	// speed := 10 // #2

	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("is it 100 mph? %t\n", speed == 100)
	fmt.Printf("is it not 100 mph? %t\n", speed != 100)

	fmt.Println(strings.Repeat("-", 25))

	// -------------------------
	// #3
	speedB := 150.5
	faster := speedB > 100 // ok: untyped

	fmt.Println("is the other car going faster?", faster)

	// ERROR: Type Mismatch
	// faster = speedB > speed
	faster = speedB > float64(speed)
	fmt.Println("is the other car going faster?", faster)

	// #4
	// ERROR:
	// only numeric values are comparable with
	// ordering operators: <, <=, >, >=

	// fmt.Println(true > faster)
}
