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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------
var string1, string2 string

func main() {
	const (
		maxTurns = 5
	)
	args := os.Args[1:]
	rand.Seed(time.Now().UnixNano())

	la := len(args)

	if la == 0 || la > 2 {
		fmt.Println("You should provide a number with or without -v for verbose output")
		return
	}

	switch la {
	case 1:
		string1 = args[0]
	case 2:
		if args[0] == "-v" {
			string2 = args[0]
			string1 = args[1]
		} else if args[1] == "-v" {
			string2 = args[1]
			string1 = args[0]
		} else {
			fmt.Println("Improper arguments were used")
		}
	}
	guess, err := strconv.Atoi(string1)
	if err != nil {
		fmt.Printf("%s not a number\n", string1)
	}
	var verbose bool
	if string2 != "" {
		verbose = true
	}

	if verbose {
		for i := 0; i <= maxTurns; i++ {
			n := rand.Intn(guess) + 1
			fmt.Printf("%d ", n)
			if guess == n {
				fmt.Printf("🎉  YOU WIN!")
				return
			}
		}
	} else {
		fmt.Printf("%d digit", guess)
		for i := 0; i <= maxTurns; i++ {
			n := rand.Intn(guess)
			if guess == n {
				fmt.Printf("🎉  YOU WIN!")
				return
			}
		}
	}
}
