// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"strconv"
	"strings"
)

const separator = ","

// Parse parses the given data and returns a slice of Properties
func parse(data string) (props []property, err error) {
	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		loc := cols[0]

		// size, err := strconv.Atoi(cols[1])
		// if err != nil {
		// 	return props, err
		// }

		// beds, err := strconv.Atoi(cols[2])
		// if err != nil {
		// 	return props, err
		// }

		// baths, err := strconv.Atoi(cols[3])
		// if err != nil {
		// 	return props, err
		// }

		// price, err := strconv.Atoi(cols[4])
		// if err != nil {
		// 	return props, err
		// }

		size, err := atoi(cols[1], err)
		beds, err := atoi(cols[2], err)
		baths, err := atoi(cols[3], err)
		price, err := atoi(cols[4], err)

		if err != nil {
			return props, err
		}

		prop := property{
			location: loc,
			size:     size,
			beds:     beds,
			baths:    baths,
			price:    price,
		}

		props = append(props, prop)
	}
	return
}

// atoi is a helper for strconv.Atoi
// it saves the previous error to simplify the error handling.
// usage:
//    n, err := atoi(p, err)
//    m, err := atoi(q, err)
func atoi(s string, err error) (int, error) {
	// if there was an error return it instead: skip the Atoi
	if err != nil {
		return 0, err
	}

	n, lerr := strconv.Atoi(s)
	if lerr != nil {
		return 0, lerr
	}
	return n, nil
}
