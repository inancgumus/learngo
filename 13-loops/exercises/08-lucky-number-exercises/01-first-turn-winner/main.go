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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	const (
		maxAtempts = 6
	)
	rand.Seed(time.Now().UnixNano())
	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf(`Not a number`)
		return
	}

	if guess < 0 {
		fmt.Printf(`Please choose a positive number`)
		return
	}
	for i := 0; i < maxAtempts; i++ {
		r := rand.Intn(10)
		fmt.Printf("%d guess number, %d pc number\n", guess, r)
		if guess == r && i == 0 {
			fmt.Println("Winner!")
			break
		}

	}
}
