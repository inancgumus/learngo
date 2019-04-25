// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Visit stores metrics for a domain
type Visit struct {
	Domain string `json:"domain"`
	Total  int    `json:"total_visits"`
	// add more metrics if needed
}

// Parser keep tracks of the parsing
type Parser struct {
	Sum     map[string]Visit `json:"visits_per_domain"` // metrics per domain
	Domains []string         `json:"unique_domains"`    // unique domain names
	Total   int              `json:"total_visits"`      // total visits for all domains
	Lines   int              `json:"-"`                 // number of parsed lines (for the error messages)
}

func main() {
	p := Parser{Sum: make(map[string]Visit)}

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.Lines++

		// Parse the fields
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, p.Lines)
			return
		}

		// Sum the total visits per domain
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], p.Lines)
			return
		}

		name := fields[0]

		// Collect the unique domains
		if _, ok := p.Sum[name]; !ok {
			p.Domains = append(p.Domains, name)
		}

		// Keep track of total and per domain visits
		p.Total += visits

		// You cannot assign to composite values
		// p.Sum[name].Total += visits

		// create and assign a new copy of `visit`
		p.Sum[name] = Visit{
			Domain: name,
			Total:  visits + p.Sum[name].Total,
		}
	}

	// Print the visits per domain
	sort.Strings(p.Domains)

	// fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	// fmt.Println(strings.Repeat("-", 45))

	// for _, name := range p.Domains {
	// 	visits := p.Sum[name]
	// 	fmt.Printf("%-30s %10d\n", name, visits.Total)
	// }

	// // Print the total visits for all domains
	// fmt.Printf("\n%-30s %10d\n", "TOTAL", p.Total)

	s, _ := json.MarshalIndent(p, "", "\t")
	fmt.Println(string(s))
}
