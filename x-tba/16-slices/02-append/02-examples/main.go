// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	var basket []string

	basket = append(basket, "banana")
	basket = append(basket, "milk", "apple")

	todo := []string{"tea", "coffee", "salt"}
	basket = append(basket, todo...)

	_ = append(basket, "pepper")

	s.Show("my basket", basket)
}
