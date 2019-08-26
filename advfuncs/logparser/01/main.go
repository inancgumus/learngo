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

/*
fmt.Println(strings.Map(unpunct, "hello!!! HOW ARE YOU???? :))"))
fmt.Println(strings.Map(unpunct, "TIME IS UP!"))

func unpunct(r rune) rune {
	if unicode.IsPunct(r) {
		return -1
	}
	return unicode.ToLower(r)
}
*/

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		r := parse(p, in.Text()) // TODO: parsed -> r
		update(p, r)
	}

	summarize(p)
	dumpErrs([]error{in.Err(), err(p)})
}
