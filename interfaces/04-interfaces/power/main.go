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
	"time"
)

// RUN IT WITH ALL THE FILES IN THE DIRECTORY LIKE SO:
// go run .

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		blender Blender
		player  Player
		kettle  Kettle
		mixer   Mixer
	)

	socket := &Socket{100}

	fmt.Printf("Socket's available power is %d kW.\n", socket.power)

	if err := socket.Plug(blender); err != nil {
		fmt.Println("Blender cannot be powered up:", err)
	}

	if err := socket.Plug(player); err != nil {
		fmt.Println("Player cannot be powered up:", err)
	}

	if err := socket.Plug(kettle); err != nil {
		fmt.Println("Kettle cannot be powered up:", err)
	}

	if err := socket.Plug(mixer); err != nil {
		fmt.Println("Mixer cannot be powered up:", err)
	}

	fmt.Printf("Socket's available power is %d kW.\n", socket.power)
}
