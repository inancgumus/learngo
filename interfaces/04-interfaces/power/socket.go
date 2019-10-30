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
	"math/rand"
)

// PowerDrawer represents an electrical device that can draw power.
// The power source doesn't have to be a Socket.
type PowerDrawer interface {
	Draw(power int)
}

// Socket has the power!
type Socket struct {
	power int
}

// Plug a device to draw power from the `Socket`
func (s *Socket) Plug(device PowerDrawer) error {
	n := rand.Intn(50) + 1

	if s.power-n < 0 {
		return fmt.Errorf("socket is out of power for %dkW", n)
	}

	s.power -= n
	device.Draw(n)

	return nil
}
