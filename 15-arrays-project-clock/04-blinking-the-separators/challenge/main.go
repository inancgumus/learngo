// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
// ★ GOAL 4: Blinking the Separators
// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

// - [ ] Blink the separators
//
//       They should be visible per two seconds.
//
//       Example: 1st second invisible
//                2nd second visible
//                3rd second invisible
//                4th second visible
//
// HINT: There are two ways to do this.
//
//   1- Manipulating the clock array directly
//      (by adding/removing the separators)
//
//   2- Deciding what to print when printing the clock

// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
// YOU CAN ALSO FIND THE FINAL SOLUTION WITH ANNOTATIONS:
// 05-full-annotated-solution
// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

package main

import (
	"fmt"
	"time"
)

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

	fmt.Print(screenClear)

	for {
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

		time.Sleep(time.Second)
	}
}
