// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		// Stores the visits per unique domain
		sum = make(map[string]int)

		// Stores the unique domain names for printing
		domains []string

		in = bufio.NewScanner(os.Stdin)
	)

	// Scan the standard-in line by line
	for line := 0; in.Scan(); line++ {

		// Parse the fields
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, line)
			return
		}
		domain, visits := fields[0], fields[1]

		// Collect the unique domains
		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		// Sum the total visits per domain
		n, _ := strconv.Atoi(visits)
		if n < 0 {
			fmt.Printf("wrong input: %q (line #%d)\n", visits, line)
			return
		}
		sum[domain] += n
	}

	// Print the visits per domain
	var total int
	for _, domain := range domains {
		hits := sum[domain]

		fmt.Printf("%-25s -> %d\n", domain, hits)
		total += hits
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-25s -> %d\n", "TOTAL", total)
}
