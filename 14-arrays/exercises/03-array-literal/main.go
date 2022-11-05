// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {
	var (
	//distances [5]int
	//data      [5]uint8
	//ratios    [1]float64
	//alives    [4]bool
	//zero      [0]uint8
	)
	names := [3]string{"Nikita", "El", "Vlad"}
	distances := [5]int{50, 40, 75, 30, 125}
	data := [5]uint8{12, 15, 11, 55, 22}
	ratios := [1]float64{3.14}
	alives := [4]bool{true, false, true, false}

	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %v\n", i, ratios[i])
	}
	for i, v := range alives {
		fmt.Printf("alives[%d]: %v\n", i, v)
	}
}
