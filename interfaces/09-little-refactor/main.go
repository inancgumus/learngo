// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func main() {
	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	l.discount(.5)
	l.print()
}

/*
Summary:

- Prefer to work directly with concrete types
  - Leads to a simple and easy to understand code
  - Abstractions (interfaces) can unnecessarily complicate your code

- Separating responsibilities is critical
  - Timestamp type can represent, store, and print a UNIX timestamp

- When a type anonymously embeds a type, it can use the methods of the embedded type as its own.
 - Timestamp embeds a time.Time
 - So you can call the methods of the time.Time through a timestamp value
*/
