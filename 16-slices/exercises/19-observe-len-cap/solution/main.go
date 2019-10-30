// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// --- #1 ---
	// var games []string
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// games := []string{}
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// games = append(games, "pacman", "mario", "tetris", "doom")
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	games := []string{"pacman", "mario", "tetris", "doom"}
	fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// --- #2 ---
	fmt.Println()

	for i := 0; i <= len(games); i++ {
		s := games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(s), cap(s))
	}

	// --- #3 ---
	fmt.Println()

	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	fmt.Println()

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	// --- #5 ---
	fmt.Println()

	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	// --- #6 ---
	fmt.Println()

	zero[0] = "command & conquer"
	games[0] = "red alert"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// uncomment and see the error.
	// _ = games[:cap(games)+1]
	// or:
	// _ = games[:5]
}
