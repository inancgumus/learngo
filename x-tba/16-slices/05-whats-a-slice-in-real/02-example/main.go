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

// Please uncomment the examples below to see the results

func main() {
	colors := []string{"black", "white"}

	// 1st example
	vivid := colors
	vivid[0] = "orange"

	// 2nd example
	// vivid = nil

	// 3th example
	vivid = append(colors, "yellow")
	vivid[0] = "gray"

	// 4th example
	colors = append(colors, "yellow")
	colors[0] = "gray"

	s.Show("colors slice header", colors)
	s.Show("vivid slice header", vivid)
}
