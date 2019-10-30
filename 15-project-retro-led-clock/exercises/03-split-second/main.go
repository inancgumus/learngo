// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Add Split Seconds
//
//  Your goal is adding the split second to the clock. A split second is
//  1/10th of a second.
//
//  1. Find the current split second
//  2. Add dot character to the clock (as in the expected output)
//  3. Add the split second digit to the clock
//  4. Blink the dot every two seconds (just like the separators)
//  5. Update the clock every 1/10th of a second, instead of every second.
//     (Update the clock every 100 millliseconds)
//
// HINTS
//   + You can find the split second using Nanosecond method of the Time type.
//     https://golang.org/pkg/time/#Time.Nanosecond
//
//   + A split second is the first digit of the Nanosecond.
//
//   + Remember: time.Second is an integer constant, so it can be divided
//               with a number.
//
//     https://golang.org/pkg/time/#Time.Second
//
// EXPECTED OUTPUT
//     Note that, clock is updated every split second instead of a second.
//
//     Separators are displayed (second is an odd number):
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       █ █
//      █    █    ░     █   █    ░    █     █       █ █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ██
//      █    █    ░   █     █    ░    █     █        █
//      █    █        ███   █         █     █        █
//      █    █    ░     █   █    ░    █     █        █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █         █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █       █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █         █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       █ █
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░     █
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █       █ █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █         █
//      █    █        ███   █         █     █         █
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░     █
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █       █ █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     Separators are not displayed (second is an even number):
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █        █     █         █   █ █       █ █
//      █    █        ███   █         █   ███       █ █
//      █    █          █   █         █   █ █       █ █
//     ███  ███       ███  ███       ███  ███       ███
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
