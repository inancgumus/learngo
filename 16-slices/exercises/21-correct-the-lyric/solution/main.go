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
	"strings"
)

func main() {
	// --- Correct Lyric ---
	// yesterday all my troubles seemed so far away
	// now it looks as though they are here to stay
	// oh I believe in yesterday
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// ------------------------------------------------------------------------
	// #1: Prepend "yesterday" to `lyric`
	// ------------------------------------------------------------------------

	//
	// --- BEFORE ---
	// all my troubles seemed so far away oh I believe in yesterday
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday
	//
	// (warning: allocates a new backing array)
	//
	lyric = append([]string{"yesterday"}, lyric...)

	// ------------------------------------------------------------------------
	// #2: Put the words to the correct positions in the `lyric` slice.
	// ------------------------------------------------------------------------

	//
	// yesterday all my troubles seemed so far away oh I believe in yesterday
	//                                              |               |
	//                                              v               v
	//                                       index: 8          pos: 13
	//
	const N, M = 8, 13

	// --- BEFORE ---
	//
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	//                                              ^
	//
	//                                              |
	//                                              +--- duplicate
	//
	// (warning: allocates a new backing array)
	lyric = append(lyric, lyric[N:M]...)

	//
	// --- BEFORE ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
	//
	// (does not allocate a new backing array because cap(lyric[:N]) > M)
	//
	lyric = append(lyric[:N], lyric[M:]...)

	// ------------------------------------------------------------------------
	// #3: Print
	// ------------------------------------------------------------------------
	fmt.Printf("%s\n", lyric)
}
