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
	p := newParser()

	fmt.Printf("%T\n", (*parser).parse)
	return

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		parsed := p.parse(in.Text())
		p.update(parsed)
	}

	p.summarize()
	dumpErrs([]error{in.Err(), p.err()})
}

// dumpErrs together to simplify multiple error handling
func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}
