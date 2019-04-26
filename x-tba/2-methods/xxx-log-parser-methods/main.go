// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	p, report := new(parser), newReport()
	for in.Scan() {
		report.update(p.parse(in.Text()))
	}

	summarize(report, p.err(), in.Err())
}
