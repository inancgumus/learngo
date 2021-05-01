// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// ages, first and last2 have the same backing arrays
	ages := []int{35, 15, 25}
	first := ages[0:1]
	last2 := ages[1:3]

	ages[0] = 55
	ages[1] = 10
	ages[2] = 20

	// grades and ages have separate backing arrays
	grades := []int{70, 99}
	grades[0] = 50

	s.Show("ages", ages)
	s.Show("ages[0:1]", first)
	s.Show("ages[1:3]", last2)
	s.Show("grades", grades)

	// let's create a new scope
	// 'cause I'm going to use variables with the same name
	{
		// ages and agesArray have the same backing arrays
		agesArray := [3]int{35, 15, 25}
		ages := agesArray[0:3]

		ages[0] = 100
		ages[2] = 50

		s.Show("agesArray", agesArray[:])
		s.Show("agesArray's ages", ages)
	}
}
