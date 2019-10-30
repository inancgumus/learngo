// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// initCells initialize the played cells to empty
func initCells() {
	// // create a string with empty cells
	// // " , , , , , , , , ,"
	// var s string

	// // number of cells = maxTurns
	// for i := 1; i <= maxTurns; i++ {
	// 	s += emptyCell // add an empty move
	// 	s += ","       // separate the cells with a comma
	// }

	// // strings are immutable — you should create a new one
	// // fortunately: most of this happens in the stack memory
	// // " , , , , , , , , ," -> // " , , , , , , , , "
	// s = strings.TrimSuffix(s, ",")

	// // store the cells in a slice (slice = list, array in other langs)
	// // Split() returns a list of strings: []string (a string slice)
	// // [" ", " ", " ", " ", " ", " ", " ", " ", " " ]
	// cells = strings.Split(s, ",")

	// Right way:
	// cells = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	//
	// Or:
	cells = make([]string, maxTurns)
	for i := range cells {
	  cells[i] = emptyCell
	}
}
