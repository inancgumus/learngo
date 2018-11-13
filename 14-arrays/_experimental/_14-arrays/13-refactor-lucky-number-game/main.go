// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game!

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	// This array will store the generated random numbers
	var pickeds [maxTurns]int

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		pickeds[turn] = n

		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			goto pickeds
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")

pickeds:
	fmt.Println("Computer has picked these:", pickeds)

	// after this line the program automatically exits
}
