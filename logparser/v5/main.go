// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"log"
	"os"

	"github.com/inancgumus/learngo/logparser/v5/pipe"
)

func main() {
	// p := pipe.Default(
	// 	os.Stdin, os.Stdout,
	// 	pipe.FilterBy(pipe.DomainExtFilter("com", "io")),
	// 	pipe.GroupBy(pipe.DomainGrouper),
	// )

	p := pipe.New(
		pipe.NewTextLog(os.Stdin),
		pipe.NewTextReport(os.Stdout),
		// pipe.NewJSONReport(os.Stdout),
		pipe.FilterBy(pipe.DomainExtFilter("com", "io")),
		pipe.GroupBy(pipe.DomainGrouper),
	)

	if err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
