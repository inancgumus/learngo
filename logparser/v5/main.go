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
	"github.com/inancgumus/learngo/logparser/v5/pipe/filter"
	"github.com/inancgumus/learngo/logparser/v5/pipe/group"
	"github.com/inancgumus/learngo/logparser/v5/pipe/parse"
	"github.com/inancgumus/learngo/logparser/v5/pipe/report"
)

func main() {
	pipe := pipe.New(
		parse.FromText(os.Stdin),
		// parse.FromJSON(os.Stdin),
		report.AsText(os.Stdout),
		filter.By(filter.Not(filter.DomainExt("com", "io"))),
		group.By(group.Domain),
		new(logger),
	)

	if err := pipe.Run(); err != nil {
		log.Fatalln(err)
	}
}

type logger struct {
	src pipe.Iterator
}

func (l *logger) Digest(records pipe.Iterator) error {
	l.src = records
	return nil
}

func (l *logger) Each(yield func(pipe.Record)) error {
	return l.src.Each(func(r pipe.Record) {
		yield(r)
	})
}
