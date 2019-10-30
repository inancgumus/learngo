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
	const (
		minsPerDay = 60 * 24
		weekDays   = 7
	)

	fmt.Printf("There are %d minutes in %d weeks.\n",
		minsPerDay*weekDays*2, 2)
}
