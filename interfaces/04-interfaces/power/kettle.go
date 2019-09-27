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

// Kettle can be powered
type Kettle struct{}

// Draw power to a Kettle
func (Kettle) Draw(power int) {
	fmt.Printf("Kettle is drawing %dkW of electrical power.\n", power)
}
