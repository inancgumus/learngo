// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"math/rand"
	"time"
)

// init is another special function
// Go calls it before the main function
func init() {
	rand.Seed(time.Now().UnixNano())
	initCells()
}

// initCells initialize the played cells to empty
func initCells() {
	for i := range cells {
		cells[i] = emptyCell
	}
}
