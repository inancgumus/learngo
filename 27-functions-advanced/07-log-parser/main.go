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
	dumpErrs(in.Err(), err(p))
}

// summarize summarizes and prints the parsing result
func summarize(p *parser) {
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	loop(p, func(r result) {
		fmt.Printf("%-30s %10d\n", r.domain, r.visits)
	})
	fmt.Printf("\n%-30s %10d\n", "TOTAL", totalVisits(p))
}

// dumpErrs simplifies handling multiple errors
func dumpErrs(errs ...error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}
