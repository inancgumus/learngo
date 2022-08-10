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
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
	const (
		max = 6
	)

	args := os.Args[1:]
	first, err1 := strconv.Atoi(args[0])
	second, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Wrong numbers")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= max; i++ {
		r := rand.Intn(10)
		fmt.Printf("%d and %d guess numbers, %d random number\n", first, second, r)
		if first == r || second == r {
			fmt.Println("You win!")
			return
		}
	}
}
