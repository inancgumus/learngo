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
	// ------------------------------------
	// #1 - THE BEST WAY
	// ------------------------------------

	students := [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sum float64

	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	const N = float64(len(students) * len(students[0]))
	fmt.Printf("Avg Grade: %g\n", sum/N)

	// ------------------------------------
	// #2 - SO SO WAY
	// ------------------------------------

	// // You don't need to define the types for the inner arrays
	// students := [2][3]float64{
	// 	[3]float64{5, 6, 1},
	// 	[3]float64{9, 8, 4},
	// }

	// var sum float64

	// sum += students[0][0] + students[0][1] + students[0][2]
	// sum += students[1][0] + students[1][1] + students[1][2]

	// const N = float64(len(students) * len(students[0]))
	// fmt.Printf("Avg Grade: %g\n", sum/N)

	// ------------------------------------
	// #3 - MANUAL WAY
	// ------------------------------------

	// student1 := [3]float64{5, 6, 1}
	// student2 := [3]float64{9, 8, 4}

	// var sum float64
	// sum += student1[0] + student1[1] + student1[2]
	// sum += student2[0] + student2[1] + student2[2]

	// const N = float64(len(student1) * 2)
	// fmt.Printf("Avg Grade: %g\n", sum/N)
}
