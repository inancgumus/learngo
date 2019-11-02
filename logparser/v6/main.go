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

	"github.com/inancgumus/learngo/logparser/v6/logly/parse"
	"github.com/inancgumus/learngo/logparser/v6/logly/record"
	"github.com/inancgumus/learngo/logparser/v6/logly/report"
)

func main() {
	var (
		p = parse.CountRecords(parse.Text(os.Stdin))
		g = record.SumGroup()
	)

	for p.Parse() {
		g.Group(p.Value())
	}
	if err := p.Err(); err != nil {
		log.Fatal(err)
	}

	if err := report.Text(os.Stdout, g.Records()); err != nil {
		log.Fatal(err)
	}
}
