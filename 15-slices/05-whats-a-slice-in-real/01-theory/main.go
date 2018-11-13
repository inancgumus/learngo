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

/*
If you find this code hard to understand, please comment
some parts of it and run it again.
*/

func main() {
	// example #1
	slice := []int{5, 6, 7, 8, 9}
	s.Show("slice", slice)

	// fmt.Println("slice[0]:", slice[0])
	// fmt.Println("slice[1]:", slice[1])
	// fmt.Println("slice[2]:", slice[2])
	// fmt.Println("slice[3]:", slice[3])
	// fmt.Println("slice[4]:", slice[4])

	// example #2
	sliced := slice[1:4]
	s.Show("slice[1:4]", sliced)

	// fmt.Println("sliced[0]:", sliced[0])
	// fmt.Println("sliced[1]:", sliced[1])
	// fmt.Println("sliced[2]:", sliced[2])
	// fmt.Println("sliced[3]:", sliced[3]) // -> you can't

	// example #3
	// the new slice will also be effected from this change
	sliced = append(sliced, 15)
	slice[1] = 200
	s.Show("append(sliced, 15)", sliced)

	// example #3b
	// the new slice won't be effected anymore
	// because, go has created a new array for the `s`
	sliced = append(sliced, 3)
	slice[1] = 0
	s.Show("slice[1] = 0", slice)
	s.Show("sliced", sliced)

	// example #4
	// its pointer will stay the same until 8 elements
	sliced = append(sliced, 10, 11, 12)
	s.Show("append(sliced, 10, 11, 12)", sliced)

	// now it will change: 13 the wicked number!
	sliced = append(sliced, 13)
	s.Show("append(sliced, 13)", sliced)

	// example #5
	var (
		// just declaring it will make it nil
		nilButHappy []int

		// without any elements will make empty
		empty = []int{}
	)

	s.Show("Empty Slice", empty)
	s.Show("Nil Slice", nilButHappy)

	fmt.Println("empty       == nil?", empty == nil)
	fmt.Println("nilButHappy == nil?", nilButHappy == nil)
}
