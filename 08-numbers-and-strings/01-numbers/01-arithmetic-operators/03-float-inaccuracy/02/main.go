// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// Go compiler sees these numbers as integers,
	//   since, there are no fractional parts in
	//   integer values,
	// So, the result becomes 1 instead of 1.5

	// So, ratio variable here is an int variable,
	//   it's because, 3 divided by 2 results
	//   in an integer.

	ratio := 3 / 2

	fmt.Printf("%d", ratio)
}
