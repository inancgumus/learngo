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

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		parsed := parse(p, in.Text())
		update(p, parsed)
	}

	summarize(p)
	dumpErrs([]error{in.Err(), err(p)})
}

// dumpErrs simplifies handling multiple errors
func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}
