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

	p := newParser()

	for in.Scan() {
		add(p, in.Text())
	}

	for _, name := range p.domains {
		// vis := p.sum[d.name]
		d := p.sum[name]

		fmt.Printf("%-25s -> %d\n", d.name, d.visits)
	}
	fmt.Printf("\n%-25s -> %d\n", "TOTAL", p.total)

	if p.lerr != nil {
		fmt.Printf("> Err: %s\n", p.lerr)
	}
}
