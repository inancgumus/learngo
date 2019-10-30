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
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// print the averages
	//
	// evil note:
	// later on, you will see how easy it would be to solve this
	// using the functions instead. there are a lot of repetitive code here.
	// -------------------------------------------------------------------------
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	// jump over the location column
	fmt.Printf("%-15s", "")

	var sum int

	for _, n := range sizes {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(sizes)))

	sum = 0
	for _, n := range beds {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(beds)))

	sum = 0
	for _, n := range baths {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(baths)))

	sum = 0
	for _, n := range prices {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(prices)))

	fmt.Println()
}
