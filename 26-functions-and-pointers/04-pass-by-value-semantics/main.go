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
		p.lines++

		d, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		p = push(p, d)
	}

	for _, d := range p.domains {
		vis := p.sum[d.name]

		fmt.Printf("%-25s -> %d\n", d.name, vis)
	}
	fmt.Printf("\n%-25s -> %d\n", "TOTAL", p.total)
}
