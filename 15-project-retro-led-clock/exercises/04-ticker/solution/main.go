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
	for shift := 0; ; shift++ {
		// we need to clear the screen here.
		// or the previous character will be left on the screen
		//
		// alternative: you can fill the rest of the missing placeholders
		//              with empty lines
		screen.Clear()
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
			l := len(clock)

			// this sets the beginning and the ending placeholder positions (indexes).
			// shift%l prevents the indexing error.
			s, e := shift%l, l

			// to slide the placeholders from the right part of the screen.
			//
			// here, we assume that as if the clock's length is double of its length.
			// this makes things easy to manage: that's why: l*2 is there.
			//
			// shift is always increasing, for it's to go beyond the clock's length,
			// it should be equal or greater than l*2, right (after the remainder of course)?
			//
			// so, if the clock goes beyond its length; this code detects that,
			// and resets the starting position to the first placeholder's index,
			// and it keeps doing so until the clock is fully displayed again.
			if shift%(l*2) >= l {
				s, e = 0, s
			}

			// print empty lines for the missing place holders.
			// this creates the effect of moving placeholders from right to left.
			//
			// l-e can only be non-zero when the above if statement runs.
			// otherwise, l-e is always zero, because l == e.
			//
			// this is one of the other benefits of assuming the length of the
			// clock as the double of its length. otherwise, l-e would always be 0.
			for j := 0; j < l-e; j++ {
				fmt.Print("     ")
			}

			// draw the digits starting from 's' to 'e'
			for i := s; i < e; i++ {
				next := clock[i][line]
				if clock[i] == colon && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
