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
		// use a comparable struct for the keys
		visits = make(map[string]int)

		domains    []string
		dupDomains = make(map[string]bool)

		line int
	)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("Wrong input: %v\n", fields)
			return
		}

		domain, svisit := fields[0], fields[1]

		dvisit, err := strconv.Atoi(svisit)
		if err != nil {
			fmt.Printf("Wrong input: %q (line #%d)\n", svisit, line)
			return
		}
		visits[domain] += dvisit

		// only add the domain once
		if !dupDomains[domain] {
			domains = append(domains, domain)
			dupDomains[domain] = true
		}
	}

	// print the summary grouped by domain
	var total int
	for _, d := range domains {
		v := visits[d]
		fmt.Printf("%-25s -> %d\n", d, v)
		total += v
	}

	fmt.Printf("\n%-25s -> %d\n", "TOTAL", total)
}
