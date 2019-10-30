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
	"strconv"
	"strings"
)

func main() {
	spendings := fetch()

	for i, daily := range spendings {
		var total int
		for _, spending := range daily {
			total += spending
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}

func fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`

	lines := strings.Split(content, "\n")

	spendings := make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			spending, _ := strconv.Atoi(field)
			spendings[i][j] = spending
		}
	}

	return spendings
}
