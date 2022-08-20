package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Variables To Variables
//
//  1. Change the value of `color` variable to "dark green"
//
//  2. Do not assign "dark green" to `color` directly
//
//     Instead, use the `color` variable again
//     while assigning to `color`
//
//  3. Print it
//
// RESTRICTIONS
//  WRONG ANSWER, DO NOT DO THIS:
//  `color = "dark green"`
//
// HINT
//  + operator can concatenate string values
//
//  Example:
//
//  "h" + "e" + "y" returns "hey"
//
// EXPECTED OUTPUT
//  dark green
// ---------------------------------------------------------

func main() {
	// UNCOMMENT THE CODE BELOW:

	color := "green"

	// ADD YOUR CODE BELOW

	color = "dark " + color

	// UNCOMMENT THE CODE BELOW TO PRINT THE VARIABLE

	fmt.Println(color)
}
