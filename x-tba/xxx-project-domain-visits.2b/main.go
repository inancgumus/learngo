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
	type visit struct {
		domain  string
		country string
	}

	var (
		// use a comparable struct for the keys
		visits = make(map[visit]int)

		domains, countries []string
		dupDomains         = make(map[string]bool)
		dupCountries       = make(map[string]bool)

		line int
	)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line++

		fields := strings.Fields(in.Text())
		if len(fields) != 3 {
			fmt.Printf("Wrong input: %v\n", fields)
			return
		}

		v := visit{
			domain:  fields[0],
			country: fields[1],
		}

		// increase the visit counts per domain + country
		fhits := fields[2]

		hits, err := strconv.Atoi(fhits)
		if err != nil {
			fmt.Printf("Wrong input: %q (line #%d)\n", fhits, line)
			return
		}
		visits[v] += hits

		// only add the domain once
		if !dupDomains[v.domain] {
			domains = append(domains, v.domain)
			dupDomains[v.domain] = true
		}

		// only add the country once
		if !dupCountries[v.country] {
			countries = append(countries, v.country)
			dupCountries[v.country] = true
		}
	}

	// print the summary grouped by domain
	for _, d := range domains {
		fmt.Println(d)

		for _, c := range countries {
			v := visit{domain: d, country: c}
			fmt.Printf("\t%s %d\n", c, visits[v])
		}
	}
}
