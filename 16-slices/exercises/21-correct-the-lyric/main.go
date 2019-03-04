package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice of lyrics data of Beatles' awesome
//  song: Yesterday. Your goal is putting the words into
//  correct positions.
//
//
// STEPS
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//  2. Put the words to the correct position in the `lyric` slice.
//
//  3. Print the `lyric` slice.
//
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
// BONUS
//
//   . Think about when does the append allocates a new backing array.
//   . Then check whether your conclusions are true or not.
//   . You can use the prettyslice package to check the backing array.
//
// ---------------------------------------------------------

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	// ...
	fmt.Printf("%s\n", lyric)
}
