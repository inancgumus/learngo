// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	speed := 100 // int
	force := 2.5 // float64

	// ERROR: invalid op
	// speed = speed * force

	// conversion can be a destructive operation
	// `force` loses its fractional part...

	speed = speed * int(force)

	fmt.Println(speed)
}
