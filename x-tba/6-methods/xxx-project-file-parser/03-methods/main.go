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
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	// create a file
	// defer closing it
	// in := bufio.NewScanner(os.Stdin)

	p := newParser()
	for in.Scan() && p.error() == nil {
		p.add(in.Text())
	}

	summarize(p)
	dumpErrs(in.Err(), p.error())
}

// funcs not always need to be reused.
// here, it tells about what it does: it summarizes the parsing result.
func summarize(p *parser) {
	// multiple iterators can be created. each one remembers the last
	// read domain record.
	next, cur := p.iterator()
	for next() {
		dom := cur()
		fmt.Printf("%-25s -> %d\n", dom.name, dom.visits)
	}
	fmt.Printf("\n%-25s -> %d\n", "TOTAL", p.total)
}

// this variadic func simplifies the multiple error handling
func dumpErrs(errs ...error) {
	for _, err := range errs {
		if err != nil {
			fmt.Printf("> Err: %s\n", err)
		}
	}
}
