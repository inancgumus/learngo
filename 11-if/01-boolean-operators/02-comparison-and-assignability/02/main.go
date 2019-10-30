// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func main() {
	var speed int
	speed = 50 // OK

	var running bool
	running = true // OK

	var force float64
	speed = int(force)

	_, _, _ = speed, running, force
}
