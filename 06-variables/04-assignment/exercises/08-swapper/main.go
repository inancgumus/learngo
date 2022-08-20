package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Swapper
//
//  1. Change `color` to "orange"
//     and `color2` to "green" at the same time
//
//     (use multiple-assignment)
//
//  2. Print the variables
//
// EXPECTED OUTPUT
//  orange green
// ---------------------------------------------------------

func main() {
	// UNCOMMENT THE CODE BELOW:

	color, color2 := "red", "blue"
	color, color2 = "orange", "green"

	fmt.Println(color, color2)
}
