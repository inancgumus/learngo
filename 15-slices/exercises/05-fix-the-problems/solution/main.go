// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	names := []string{"Einstein", "Shepard", "Tesla"}

	books := []string{"Stay Golden", "Fire", "Kafka's Revenge"}
	sort.Strings(books)

	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}
	sort.Ints(nums[:])

	fmt.Printf("%q\n", strings.Join(names, " and "))
	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
