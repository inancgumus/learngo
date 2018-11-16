// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// we've received the raining probabilities
	data := []float64{10, 25, 30, 50}

	s.Show("Probabilities", data)

	fmt.Printf("Is it gonna rain? %.f%% chance.\n",
		(data[0]+data[1]+data[2]+data[3])/
			float64(len(data)))

	// -----------------------------------------------------
	// but it turns out that the first two items of the data
	// has been corrupted by a hacker.
	//
	// this time, we've received clean data.
	// let's overwrite the invalid data by copying

	copy(data, []float64{80, 90})

	// why copy? why not just assign and overwrite?
	// because: it overwrites the whole slice
	// data = []float64{80, 90}

	s.Show("Probabilities", data)

	fmt.Printf("Is it gonna rain? %.f%% chance.\n",
		(data[0]+data[1]+data[2]+data[3])/
			float64(len(data)))
}
