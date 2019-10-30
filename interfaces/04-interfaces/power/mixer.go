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

// Mixer can be powered
type Mixer struct{}

// Draw power to a Mixer
func (Mixer) Draw(power int) {
	fmt.Printf("Mixer is drawing %dkW of electrical power.\n", power)
}
