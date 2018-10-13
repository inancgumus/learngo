// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// order of conversion matters...

func main() {
	speed := 100
	force := 2.5

	fmt.Printf("speed     : %T\n", speed)
	fmt.Printf("conversion: %T\n", float64(speed))
	fmt.Printf("expression: %T\n", float64(speed)*force)

	// TYPE MISMATCH:
	//   speed is int
	//   expression is float64
	// speed = float64(speed) * force

	// CORRECT: int * int
	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
