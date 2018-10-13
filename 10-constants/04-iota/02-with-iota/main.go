// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
}
