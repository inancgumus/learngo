package main

import (
	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Add newlines to the lyric sentences
//
//  You have a slice of lyrics data of Beatles' awesome
//  song: Yesterday. Your goal is adding newlines after
//  each sentence of the lyric.
//
//  Your goal is creating a new slice, then copying
//  each sentence of the lyric to the new buffer, then
//  after each sentence adding a newline character.
//
//  You cannot guess how many times you will get across
//  something like this. Believe me, learning this will
//  make you a better Gopher.
//
//
// RESTRICTIONS
//
//  . For warming-up, in this exercise, never use the `append()` func.
//
//  	. Instead, only use the `copy()` func.
//
//  . You cannot use slicing for printing the sentences.
//
//    . Instead, use slicing to fill the new buffer, then print it.
//
//
// STEPS
//
//  . Follow the instructions inside the code.
//
//
// EXPECTED OUTPUT
//
//  yesterday all my troubles seemed so far away
//  now it looks as though they are here to stay
//  oh i believe in yesterday
//
// ---------------------------------------------------------

func main() {
	//
	// YOU DON'T NEED TO TOUCH THIS
	//
	// This inits some options for the prettyslice package.
	// You can change the options if you want.
	//
	// s.Colors(false)     // if your editor is light colored then enable this
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 15      // prints max 15 elements per line
	s.SpaceCharacter = "*" // print this instead of printing a newline (for debugging)

	//
	// UNCOMMENT THE VARIABLE BELOW THEN START!
	//

	// lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday`)

	//
	// RESTRICTION EXPLANATION:
	//
	// Don't do something like this:
	// (Do not use slicing for printing the sentences)
	//

	// fmt.Println(lyric[:8])
	// fmt.Println(lyric[8:18])
	// fmt.Println(lyric[18:23])

	// ========================================================================
	//
	// #1: CREATE A LARGE ENOUGH BUFFER (A NEW SLICE)
	//
	//     You need to put each lyric sentence + a newline into this buffer
	//

	// I name the buffer: `fix`, you can name it however you want
	// fix := make(...)

	// ========================================================================
	//
	// #2: CALCULATE THE CUT POINTS
	//
	//     You want to put newlines after each sentence. So you may want to put
	//     these index positions into a slice. Then you can use them to cut the
	//     lyric slice.
	//

	// cutpoints := []int{...}

	// ========================================================================
	//
	// #3: CREATE A LOOP AND COPY THE SENTENCES INTO THE BUFFER
	//

	// for ... {
	//   Use the `copy` function to copy from the `lyric` slice to the buffer.
	//
	//   Copy a newline character (in a string) to the buffer after each sentence.
	//
	//   You can use slicing here for filling the new buffer.
	//
	//   Uncomment this to aid debugging (to see how the fix slice changes)
	//   s.Show("fix slice", fix)
	// }

	// ========================================================================
	//
	// #4: PRINT THE BUFFER
	//

	// ...
}
