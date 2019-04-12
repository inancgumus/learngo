// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pricing

import (
	"strconv"
	"strings"
)

// named result values
func parse(data string) (props []Property) {
	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		size, _ := strconv.Atoi(cols[1])
		beds, _ := strconv.Atoi(cols[2])
		baths, _ := strconv.Atoi(cols[3])
		price, _ := strconv.Atoi(cols[4])

		prop := Property{
			Location: cols[0],
			Size:     size,
			Beds:     beds,
			Baths:    baths,
			Price:    price,
		}

		props = append(props, prop)
	}
	return
}
