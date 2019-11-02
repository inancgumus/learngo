// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"log"
	"os"
)

func main() {
	// p := pipe.Default(
	// 	os.Stdin, os.Stdout,
	// 	pipe.FilterBy(pipe.DomainExtFilter("com", "io")),
	// 	pipe.GroupBy(pipe.DomainGrouper),
	// )

	p, err := fromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
