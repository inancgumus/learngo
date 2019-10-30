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
	"os"

	"github.com/inancgumus/learngo/logparser/testing/report"
)

func main() {
	p := report.New()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.Parse(in.Text())
	}

	summarize(p.Summarize(), p.Err(), in.Err())
}
