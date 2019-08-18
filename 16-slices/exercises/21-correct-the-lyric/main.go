package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
// STEPS
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//  3. Print the `lyric` slice.
//
//     HINT: You don't need use the "yesterday" word from the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocates a new backing array.
//
//   + Check whether your conclusions are correct.
//
// ---------------------------------------------------------

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	// ...
	fmt.Printf("%s\n", lyric)
}
