// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	var (
		speed     = 100
		prevSpeed = 50
	)

	speed, prevSpeed = prevSpeed, speed

	fmt.Println(speed, prevSpeed)
}
