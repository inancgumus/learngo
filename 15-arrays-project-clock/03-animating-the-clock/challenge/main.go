// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
// ★ GOAL 3  : Animate the Clock
// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

// - [ ] Create an infinite loop to update the clock
//
//
// - [ ] Update the clock every second
//
//       time.Sleep(time.Second) will stop the world for 1 second
//
//       See this for more info:
//       https://golang.org/pkg/time/#Sleep
//
//
// - [ ] Clear the screen before the infinite loop
//
//       + Get my library for clearing the screen:
//
//           go get -u github.com/inancgumus/screen
//
//       + Then, import it and call it in your code like this:
//
//           screen.Clear()
//
//       + If you're using Go Playground instead, do this:
//
//           fmt.Println("\f")
//
//
// - [ ] Move the cursor to the top-left corner of the screen,
//       before each step of the infinite loop
//
//       + Call this in your code like this:
//
//           screen.MoveTopLeft()
//
//       + If you're using Go Playground instead, do this again:
//
//           fmt.Println("\f")
//
//
// -----------------------------------------------------------------------------
// SIDE NOTE FOR THE CURIOUS
// -----------------------------------------------------------------------------
//
// If you're curious about how my screen clearing package works, read on.
//
//   On bash, it uses special commands, if you open the code, you can see that.
//
//     \033 is a special control code:
//     [2J clears the screen and the cursor
//     [H  moves the cursor to 0, 0 screen position
//
//     See for more info:
//     https://bluesock.org/~willkg/dev/ansi.html
//
//   On Windows, I'm directly calling the Windows API functions. This is way
//   advanced at this stage of the course, however, I'll explain it afterward.
//
//   My package automatically adjusts itself depending on where it is compiled.
//   On Windows, it uses the special Windows API calls;
//   On other operating systems, it uses the bash special commands
//   that I've explained above.

package main

import (
	"fmt"
	"time"
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

	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

	// [8][5]string
	clock := [...]placeholder{
		// extract the digits: 17 becomes, 1 and 7 respectively
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
}
