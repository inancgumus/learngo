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
)

func main() {
	var (
		speed int
		now   time.Time
	)

	speed, now = 100, time.Now()

	fmt.Println(speed, now)

	// EXERCISE:
	//   Try this alternative (formatted time).

	// fmt.Println(speed, now.Format(time.Kitchen))
}
