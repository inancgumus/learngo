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

// Player can be powered
type Player struct{}

// Draw power to a Player
func (Player) Draw(power int) {
	fmt.Printf("Player is drawing %dkW of electrical power.\n", power)
}
