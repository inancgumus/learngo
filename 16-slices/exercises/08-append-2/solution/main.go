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
	"time"
)

func main() {
	// ------------------------------------------------------------
	// Create nil slices
	// ------------------------------------------------------------

	// Pizza toppings
	var pizza []string

	// Departure times
	var departures []time.Time

	// Student graduation years
	var grads []int

	// On/off states of lights in a room
	var lights []bool

	// ------------------------------------------------------------
	// Append them some elements (use your creativity!)
	// ------------------------------------------------------------
	pizza = append(pizza, "pepperoni", "onions", "extra cheese")

	now := time.Now()
	departures = append(departures,
		now,
		now.Add(time.Hour*24), // 24 hours after `now`
		now.Add(time.Hour*48)) // 48 hours after `now`

	grads = append(grads, 1998, 2005, 2018)

	lights = append(lights, true, false, true)

	// ------------------------------------------------------------
	// Print the slices
	// ------------------------------------------------------------

	fmt.Printf("pizza       : %s\n", pizza)
	fmt.Printf("\ngraduations : %d\n", grads)
	fmt.Printf("\ndepartures  : %s\n", departures)
	fmt.Printf("\nlights      : %t\n", lights)
}
