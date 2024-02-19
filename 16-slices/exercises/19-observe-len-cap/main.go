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
)

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// var games []string
	// fmt.Printf("length: %d ; capacity : %d\n", len(games), cap(games))
	// games := []string{}
	// fmt.Printf("length: %d ; capacity : %d\n", len(games), cap(games))
	// games = append(games, "pacman", "mario", "tetris", "doom")
	// fmt.Printf("length: %d ; capacity : %d\n", len(games), cap(games))

	games := []string{"pacman", "mario", "tetris", "doom"}

	// --- #2 ---
	fmt.Println()
	for i := 0; i <= 4; i++ {
		games = games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(games), cap(games))
	}

	// --- #3 ---
	// 2. print the games and the new slice's len and cap
	//
	// 3. append a new element to the new slice
	//
	// 4. print the new slice's lens and caps
	//
	// 5. repeat the last two steps 5 times (use a loop)
	//
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	// []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	fmt.Println()

	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	for _, game := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, game)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	fmt.Println()

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	zero = zero[:cap(zero)]
	for n := range zero {
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(zero[:n]), cap(zero[:n]), zero[:n])
	}

	// --- #6 ---
	// 4. observe that they don't have the same backing array
	fmt.Println()
	zero[1] = "zero zelda"
	games[1] = "zelda"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	_ = games[:5]
}
