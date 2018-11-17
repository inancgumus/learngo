// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"time"
)

// GOALS:
// 1. Update the clock in a loop
//
// 2. Before the whole loop:
//    Clear the screen once
//
// 3. Before each step of the loop:
//    Move the cursor to top-left position

// So, how to clear the screen and move the cursor?
// Check out the following constants.
//
// screenClear  : When printed clears the screen.
// cursorMoveTop: Moves the cursor to top-left screen position
//
// + Note: This only works on bash command prompts.
//
// + For Go Playground, use these instead:
//   screenClear   = "\f"
//   cursorMoveTop = "\f"
//
// See for more info: https://bluesock.org/~willkg/dev/ansi.html
const (
	screenClear   = "\033[2J"
	cursorMoveTop = "\033[H"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// clears the command screen
	fmt.Print(screenClear)

	// Go Playground will not run an infinite loop.
	// So, instead, you may loop for 1000 times:
	// for i := 0; i < 1000; i++ {
	for {
		// moves the cursor to the top-left position
		fmt.Print(cursorMoveTop)

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
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}
}
