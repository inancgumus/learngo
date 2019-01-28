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
	var todo []string

	todo = append(todo, "sing")

	// you can only append elements with the same element type of the slice
	// todo = append(todo, 42)

	todo = append(todo, "run")

	// append is a variadic function, so you can append multiple elements
	todo = append(todo, "code", "play")

	// you can also append a slice to another slice using ellipsis: ...
	tomorrow := []string{"see mom", "learn go"}
	todo = append(todo, tomorrow...)
	// todo = append(todo, "see mom", "learn go")

	s.Show("todo", todo)
}
