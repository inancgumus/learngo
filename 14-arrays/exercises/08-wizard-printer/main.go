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
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {
	first_name := [...]string{
		"Albert",
		"Isaac",
		"Stephen",
		"Marie",
		"Charles",
	}
	last_name := [...]string{
		"Einstein",
		"Newton",
		"Hawking",
		"Curie",
		"Darwin",
	}

	nickname := [...]string{
		"time",
		"apple",
		"blackhole",
		"radium",
		"fittest",
	}
	scientists := [3][5]string{
		first_name,
		last_name,
		nickname,
	}
	fmt.Printf("| %-15s | %-15s | %-15s |\n", "name", "surname", "nickname")
	for i := range scientists[0] {
		for j := range scientists {
			fmt.Printf("| %-15s ", scientists[j][i])
			if j == len(scientists)-1 {
				fmt.Println("|\r")
			}
		}
	}

}
