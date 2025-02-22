// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main
import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable

	// #3
	// Key        : Last name
	// Element    : Phone numbers

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	phonenumber:=map[string]int{
		"gupta":9412366123,
		"aggarwal":9458840400,
		"tayal":9412580180,
		"mittal":9412200220,
		}
	product:=map[string]int{
		"pen":1,
		"ball":2,
		"pencil":3,
		"eraser":4
		}
	list:=map[string][]int{
		"gupta":[]int{9412366123,9412104321,9412580180},
		"mittal":[]int{9412200220,9458840400},
		}
	basket:=map[int]map[int]int{
	345:map[int]int{
		1:12,
		2:4,
		3:1,
		4:1,
		       },
	567:map[int]int{
		56:1,
		45:3,
		},
	32:map[int]int{
		44:12,
		},
  }
	
		
}
