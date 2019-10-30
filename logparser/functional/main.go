// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

func main() {
	p := pipeline{
		read:   textReader(os.Stdin),
		write:  textWriter(os.Stdout),
		filter: notUsing(domainExtFilter("io", "com")),
		group:  domainGrouper,
	}

	// var p pipeline
	// p.
	// 	filterBy(notUsing(domainExtFilter("io", "com"))).
	// 	groupBy(domainGrouper)

	if err := p.start(); err != nil {
		fmt.Println("> Err:", err)
	}
}

// []outputter{textFile("results.txt"), chartFile("graph.png")}

// func outputs(w io.Writer) outputFn {
// 	tw := textWriter(w)
// 	cw := chartWriter(w)

// 	return func(rs []result) error {
// 		err := tw(rs)
// 		err = cw(rs)
// 		return err
// 	}
// }
