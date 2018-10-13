// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	var speed int
	// speed = "100" // NOT OK

	var running bool
	// running = 50 // NOT OK

	var force float64
	// speed = force // NOT OK

	_, _, _ = speed, running, force
}
