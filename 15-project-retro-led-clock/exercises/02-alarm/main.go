// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Set an Alarm
//
//  Goal is printing " ALARM! " every 10 seconds.
//
// EXPECTED OUTPUT
//
//     ██   ███       ███  ██        ███  ███
//      █     █   ░     █   █    ░     █  █ █
//      █   ███       ███   █        ███  ███
//      █   █     ░     █   █    ░     █    █
//     ███  ███       ███  ███       ███  ███
//
//          ███  █    ███  ██   █ █   █
//          █ █  █    █ █  █ █  ███   █
//          ███  █    ███  ██   █ █   █
//          █ █  █    █ █  █ █  █ █
//          █ █  ███  █ █  █ █  █ █   █
//
//     ██   ███       ███  ██        █ █  ██
//      █     █   ░     █   █    ░   █ █   █
//      █   ███       ███   █        ███   █
//      █   █     ░     █   █    ░     █   █
//     ███  ███       ███  ███         █  ███
//
// HINTS
//
//  <<< BEWARE THE SPOILER! >>>
//
//  I recommend you to first solve the exercise yourself before reading the
//  following hint.
//
//
//  + You can create a new array named alarm with the same length of the
//    clocks array, so you can copy the alarm array to the clocks array
//    every 10 seconds.
//
// ---------------------------------------------------------

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				// colon blink
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
