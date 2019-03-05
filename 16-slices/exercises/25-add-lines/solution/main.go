// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

//
// ! NOTE If the program does not work, please update your
//        local copy of the prettyslice package:
//
//        go get -u github.com/inancgumus/prettyslice
//

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
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)

	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday`)

	// CREATE A LARGE ENOUGH BUFFER
	// `+3` because we're going to add 3 newline characters
	fix := make([]string, len(lyric)+3)

	// CALCULATE THE CUT POINTS
	//
	// + The first sentence has 8 words so its cutpoint is 8.
	//
	// + The second one has 10 words so its cutpoint is 10.
	//
	// + The third one has 5 words so its cutpoint is 5.
	//
	//
	// yesterday all my troubles seemed so far away now it looks as though they are here to stay
	//                                         |
	//                                         v
	//                               cutpoint: 8
	//
	// ... now it looks as though they are here to stay oh i believe in yesterday
	//                                             |
	//                                             v
	//                                             10
	//
	// ... now it looks as though they are here to stay oh i believe in yesterday
	//                                                                  |
	//                                                                  v
	//                                                                  5
	cutpoints := []int{8, 10, 5}

	// `n` tracks how much we've moved inside the `lyric` slice
	// `i` tracks the sentence that we're on
	for i, n := 0, 0; n < len(lyric); i++ {
		//
		// copy to `fix` from the `lyric`
		//
		//   destination:
		//     fix[n+i] because we don't want to delete the previous copy.
		//     it moves sentence by sentence, using the cutpoints.
		//
		//   source:
		//     lyric[n:n+cutpoints[i]] because we want copy the next sentence
		//     beginning from the number of the last copied elements to the
		//     n+next cutpoint (the next sentence).
		//
		n += copy(fix[n+i:], lyric[n:n+cutpoints[i]])

		// add a newline after the number of copied elements
		// notice that the '\n' position slides as we move over
		// that's why n+i
		fix[n+i] = "\n"

		// uncomment this to aid debugging (to see how the fix slice changes)
		// s.Show("fix slice", fix)
	}

	// print the fix slice
	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}
