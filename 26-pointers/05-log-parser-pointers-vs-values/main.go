// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		parsed := parse(p, in.Text())
		update(p, parsed)
	}

	summarize(p)
	dumpErrs([]error{in.Err(), err(p)})
}

// summarize summarizes and prints the parsing result
func summarize(p *parser) {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}

// dumpErrs simplifies handling multiple errors
func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}
