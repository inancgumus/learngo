// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	const (
		EST = -(5 + iota) // CORRECT: -5
		_
		MST // CORRECT: -7
		PST // CORRECT: -8
	)

	fmt.Println(EST, MST, PST)
}
