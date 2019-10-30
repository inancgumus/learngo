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
	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	l.discount(.5)

	// The list is a stringer.
	// The `fmt.Print` function can print the `l`
	// by calling `l`'s `String()` method.
	//
	// Underneath, `fmt.Print` uses a type switch to
	// detect whether a type is a Stringer:
	// https://golang.org/src/fmt/print.go#L627
	fmt.Print(l)

	// The money type is a stringer.
	// You don't need to call the String method when printing a value of it.
	// var pocket money = 10
	// fmt.Println("I have", pocket)
}

/*
Summary:

- fmt.Stringer has one method: String()
  - That returns a string.
  - It is better to be an fmt.Stringer instead of printing directly.

- Implement the String() on a type and the type can represent itself as a string.
  - Bonus: The functions in the fmt package can print your type.
  - They use type assertion to detect if a type implements a String() method.

- strings.Builder can efficiently combine multiple string values.
*/
