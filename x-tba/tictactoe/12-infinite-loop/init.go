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
