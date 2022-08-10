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

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

func main() {
	const (
		maxAttempts = 6
	)
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Wrong number!")
	}
	for i := 0; i <= maxAttempts; i++ {
		r := rand.Intn(10)
		fmt.Printf("%d guess nunberm %d pc number\n", guess, r)
		var (
			looser   string
			greeting string
		)
		switch {
		case r <= 5:
			greeting = "YOU WON"
			looser = "LOSER!"
		case r > 5:
			greeting = "YOU AWESOME"
			looser = "YOU LOST. TRY AGAIN?"

		}
		if guess == r {
			fmt.Printf(greeting + "\n")
			return
		} else {
			fmt.Printf(looser + "\n")
		}
	}

}
