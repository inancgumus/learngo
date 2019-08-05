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

	"github.com/inancgumus/learngo/29-methods/logparser-pkg/report"
)

func main() {
	p := report.New()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.Parse(in.Text())
	}

	summarize(p.Summarize(), p.Err(), in.Err())
}
