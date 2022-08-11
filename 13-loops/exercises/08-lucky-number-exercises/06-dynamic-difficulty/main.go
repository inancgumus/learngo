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
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game too easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		maxTurns = 5
		delta    = 0
	)
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Incorrect number of args")
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Not a number")
		return
	}
	if guess > 10 {
		delta = guess / 4
	}

	for turn := maxTurns + delta; turn > 0; turn-- {
		n := rand.Intn(guess) + 1
		fmt.Printf("%d ", n)
		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}
	fmt.Println("☠️  YOU LOST... Try again?")
}
