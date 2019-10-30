// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

func setGameSpeed() {
	// args can be used within the if block
	// gs and err can be used in the else if and else branches
	if args := os.Args; len(args) == 1 {
		fmt.Println("Setting game speed to default.")
	} else if gs, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Wrong game speed. Provide an integer.")
	} else {
		gameSpeed = time.Duration(gs) * time.Second
	}

	fmt.Printf("Game speed is %q.\n", gameSpeed)
}
