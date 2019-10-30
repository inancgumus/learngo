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

// Blender can be powered
type Blender struct{}

// Draw power to a Blender
func (Blender) Draw(power int) {
	fmt.Printf("Blender is drawing %dkW of electrical power.\n", power)
}
