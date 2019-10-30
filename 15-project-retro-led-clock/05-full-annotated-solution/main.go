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
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// for keeping things easy to read and type-safe
	type placeholder [5]string

	// put the digits (placeholders) into variables
	// using the placeholder array type above
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

	// This array's type is "like": [10][5]string
	//
	// However:
	// + "placeholder" is not equal to [5]string in type-wise.
	// + Because: "placeholder" is a defined type, which is different
	//   from [5]string type.
	// + [5]string is an unnamed type.
	// + placeholder is a named type.
	// + The underlying type of [5]string and placeholder is the same:
	//     [5]string
	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// For Go Playground, do not use this.
	screen.Clear()

	// Go Playground will not run an infinite loop.
	// Loop for example 1000 times instead, like this:
	//   for i := 0; i < 1000; i++ { ... }
	for {
		// For Go Playground, use this instead:
		// fmt.Print("\f")
		screen.MoveTopLeft()

		// get the current hour, minute and second
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		// extract the digits: 17 becomes, 1 and 7 respectively
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		// Explanation: clock[0]
		// + Each element of clock has the same length.
		// + So: Getting the length of only one element is OK.
		// + This could be: "zero" or "one" and so on... Instead of: digits[0]
		//
		// The range clause below is ~equal to the following code:
		//   line := 0; line < len(clock[0]); line++
		for line := range clock[0] {
			// Print a line for each placeholder in clock
			for index, digit := range clock {
				// Colon blink on every two seconds.
				// + On each sec divisible by two, prints an empty line
				// + Otherwise: prints the current pixel
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				// Print the next line and,
				// give it enough space for the next placeholder
				fmt.Print(next, "  ")
			}

			// After each line of a placeholder, print a newline
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}
}
