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
	basket := []string{"banana", "apple", "coffee"}
	s.Show("basket (before)", basket)

	// #1 example: type
	fmt.Printf("[]string type: %T\n", basket)

	// #2 example: getting and setting elements
	fmt.Println(basket[0], basket[1], basket[2])

	basket[0] = "pepper"
	basket[1] = "water"
	basket[2] = "tea"
	s.Show("basket (after)", basket)

	// #3 example: empty slice
	emptyBasket := []rune{}
	s.Show(`emptyBasket := []rune{}`, emptyBasket)

	// #4 example: nil slice
	// a slice's zero value is nil
	var nilButHappy []rune

	s.Show(`var nilButHappy []rune`, nilButHappy)

	// #5 example: comparing to nil
	fmt.Println("nilButHappy == nil", nilButHappy == nil)
	fmt.Println("emptyBasket == nil", emptyBasket == nil)

	// you can't compare slices other than nil
	// nilButHappy == emptyBasket

	// #6 example: comparing slices
	newBasket := []string{"pepper", "water", "tea"}

	equal := true
	for i := range basket {
		if basket[i] != newBasket[i] {
			equal = false
			break
		}
	}

	s.Show("equal?", basket, newBasket)
	fmt.Println(equal)
}
