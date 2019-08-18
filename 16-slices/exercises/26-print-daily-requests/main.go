package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got request logs of a web server. The log data
//  contains 8-hourly totals per each day. Find and print
//  the total requests per day, as well as the grand total.
//
//  See the `reqs` slice and the steps in the code below.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	// ================================================
	// Don't touch this code.

	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, 100, // 4th day: 1100 requests
		// grand total: 5400 requests
	}
	// ================================================

	_ = reqs // remove this when you start

	// ================================================
	// Allocate a slice efficiently with the exact size needed.
	//
	// Find the `size` automatically using the `reqs` slice.
	// Do not type it manually.
	//
	// Change this code:
	size := 0 // you need to find the size.
	daily := make([][]int, 0, size)
	//
	// ================================================

	// ================================================
	// Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]
	//
	// Change this code:
	//
	// for N < len(reqs) {
	//   append the daily requests to `daily` per 8-hour groups
	// }
	// ================================================

	// ================================================
	// Don't touch the following code:

	// Print the header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Print the data per day along with the totals
	var grand int

	for i, day := range daily {
		var sum int

		for _, req := range day {
			sum += req
			fmt.Printf("%-10d%-10d\n", i+1, req)
		}

		fmt.Printf("%9s %-10d\n\n", "TOTAL:", sum)

		grand += sum
	}
	fmt.Printf("%9s %-10d\n", "GRAND:", grand)

	// ================================================
}
