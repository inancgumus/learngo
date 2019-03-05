package main

import (
	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Practice advanced slice operations
//
//  This exercise's intention is warm you up and reinforce
//  your memory for the advanced slice operations.
//
//  Please follow the directions in the following code.
//
//  To see the expected output, please run:
//
//    go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	//
	// Use the prettyslice package for printing the slices,
	// or do it with your own Printf that matches the output
	// of the prettyslice package.
	//

	// This allows you to see the backing array of the slices.
	s.PrintBacking = true
	// Shows 10 slice elements per line
	s.MaxPerLine = 10
	// Prints 60 character per line
	s.Width = 60

	// ########################################################
	//
	// #1: Create a string slice with a length and capacity
	//     of 5. Call the slice: `names` and print it.
	//

	// ...
	// s.Show("1st step", names)

	// ########################################################
	//
	// #2: Append the following names to the names slice:
	//
	//     "einstein", "tesla", "aristo"
	//
	//     Print the names slice.
	//
	//     Observe how the slice and its backing array change.
	//

	// ...
	// s.Show("2nd step", names)

	// ########################################################
	//
	// #3: Fix the append problem by reinitializing the
	//     names slice below (currently the previous code
	//     appends after 5 elements).
	//
	//     Append the new elements to the head of the names
	//     slice instead.
	//
	//     Print the names slice.
	//
	// HINT:
	//
	//    The problem is in the `make` function.
	//

	// ...
	// s.Show("3rd step", names)

	// ########################################################
	//
	// #4: Copy elements from an array to the `names` slice.
	//
	//     Copy only the first two elements of the following
	//     array to the last two elements of the `names` slice.
	//
	//     Print the names slice (do not forget extending it
	//     after appending). You should print 5 elements.
	//
	//     Observe how the backing array stays the same.
	//

	// Array (uncomment):
	// moreNames := [...]string{"plato", "khayyam", "ptolemy"}

	// ...

	// s.Show("4th step", names)

	// ########################################################
	//
	// #5: Clone the `names` slice to the `clone` slice.
	//
	//     But only clone (copy) the last 3 elements of the
	//     `names` slice.
	//
	//     Then, append the first two elements of the `names`
	//     slice to the `clone` slice.
	//
	//     Ensure that after appending no new backing array
	//     allocations occur for the `clone` slice.
	//
	//     Print the clone slice before and after the append.
	//

	// ...
	// s.Show("5th step (before append)", clone)

	// ...
	// s.Show("5th step (after append)", clone)

	// ########################################################
	//
	// #6: Slice the `clone` slice between 2nd and 4th
	//     elements.
	//
	//     Put the sliced slice into a variable named:
	//     `sliced`.
	//
	//     Append "hypatia" to the `sliced` slice.
	//
	//     Ensure that new backing array allocation "happens".
	//
	//       Change the 3rd element of the `clone` slice
	//       to "elder".
	//
	//       Doing so should not change any elements of
	//       the `sliced` slice.
	//
	//     Print the `clone` and `sliced` slices.
	//

	// ...
	// s.Show("6th step", clone, sliced)
}
