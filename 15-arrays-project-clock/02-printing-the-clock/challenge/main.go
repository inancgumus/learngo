// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
// ★ GOAL 2  : Printing the Clock
// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
//
// - [ ] Get the current time
//
// - [ ] Get the current hour, minute and second from the current time
//
// - [ ] Create the clock array
//
//   - [ ] Get the individual digits from the digits array
//
// - [ ] Print the clock
//
//   - [ ] In the loops, use the clocks array instead
//
// - [ ] Create a separator array (it's also a placeholder)
//
// - [ ] Add the separators into the correct positions of
//       the clock array
//
// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
// ★ Usable Artifacts
// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

// Clock Characters:
//
//   You can put these in constants if you like.
//
//   Use this for the digits       : "█"
//   Use this for the separators   : "░"

package main

import (
	"fmt"
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

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}
}
