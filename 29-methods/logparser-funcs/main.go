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
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		res := parse(p, in.Text())
		updateSummary(p.summary, res)
	}

	summarize(summarizeParse(p))
	dumpErrs(errParse(p), in.Err())
}
