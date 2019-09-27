// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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
