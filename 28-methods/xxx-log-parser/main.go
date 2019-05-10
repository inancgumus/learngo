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

	"github.com/inancgumus/learngo/28-methods/xxx-log-parser/report"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	r := report.New()
	for in.Scan() {
		r.Parse(in.Text())
	}

	summarize(r.Summarize(), r.Err(), in.Err())
}
