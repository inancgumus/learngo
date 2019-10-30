// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// VARIABLE NAMING RULES

func main() {
	// CORRECT DECLARATIONS
	var speed int
	var SpeeD int

	// underscore is allowed but not recommended
	var _speed int

	// Unicode letters are allowed
	var 速度 int

	// keep the compiler happy
	_, _, _, _ = speed, SpeeD, _speed, 速度
}
