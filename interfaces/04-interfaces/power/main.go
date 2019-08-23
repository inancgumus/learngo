// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Socket has the power!
type Socket struct {
	power int
}

// Plug a device to draw power from the `Socket`
func (s *Socket) Plug(device PowerDrawer) {
	n := rand.Intn(50)
	s.power -= n
	device.Draw(n)
}

// PowerDrawer can be any type that can be powered
type PowerDrawer interface {
	Draw(power int)
}

// -----

// Kettle can be powered
type Kettle struct{}

// Draw power to a Kettle
func (Kettle) Draw(power int) {
	fmt.Printf("Kettle is drawing %dkW of electrical power.\n", power)
}

// Mixer can be powered
type Mixer struct{}

// Draw power to a Mixer
func (Mixer) Draw(power int) {
	fmt.Printf("Mixer is drawing %dkW of electrical power.\n", power)
}

// -----

func main() {
	rand.Seed(time.Now().UnixNano())

	kettle := Kettle{}
	mixer := Mixer{}

	socket := &Socket{100}
	socket.Plug(kettle)
	socket.Plug(mixer)

	fmt.Printf("Socket's available power is %d kW.\n", socket.power)
}
