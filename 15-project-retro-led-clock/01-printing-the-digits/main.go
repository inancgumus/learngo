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
	"strings"
)

func main() {
	digits := [...][5]string{
		{"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		},
		{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		}, {
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		}, {
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		}, {
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		},
	}

	for line_number := range digits[0] {
		for _, digit := range digits {
			fmt.Print(digit[line_number], strings.Repeat(" ", 2))
		}
		fmt.Println()
	}
}
