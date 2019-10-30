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
	type Duration int64

	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)

	// The types are main.Duration instead of just Duration
	// It's because, Duration type now belongs to main func
	fmt.Printf("Nanosecond is %T\n", Nanosecond)
	fmt.Printf("Microsecond is %T\n", Microsecond)
	fmt.Printf("Millisecond is %T\n", Millisecond)
	fmt.Printf("Second is %T\n", Second)
	fmt.Printf("Minute is %T\n", Minute)
	fmt.Printf("Hour is %T\n", Hour)

	// Print the types of time.Duration constants directly
	// This time, types will be time.Duration
	// It's because, they're inside the time package
	fmt.Printf("Nanosecond is %T\n", time.Nanosecond)
	fmt.Printf("Microsecond is %T\n", time.Microsecond)
	fmt.Printf("Millisecond is %T\n", time.Millisecond)
	fmt.Printf("Second is %T\n", time.Second)
	fmt.Printf("Minute is %T\n", time.Minute)
	fmt.Printf("Hour is %T\n", time.Hour)
}
