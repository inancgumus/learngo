// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/inancgumus/learngo/logparser/v6/parse"
	"github.com/inancgumus/learngo/logparser/v6/report"
)

func main() {
	// trace.Start(os.Stderr)
	// defer trace.Stop()

	var p parse.Parser
	// p = parse.Text(os.Stdin)
	p = parse.JSON(os.Stdin)
	p = parse.CountRecords(p)

	r := report.Text(os.Stdout)

	var out []parse.Record
	for p.Parse() {
		r := p.Value()

		// if !parse.Filter(r) {
		// 	continue
		// }

		// sum.group(r)
		out = append(out, r)
	}

	if err := p.Err(); err != nil {
		log.Fatalln(err)
	}

	// var out []parse.Record
	// for sum.More() {

	// }

	if err := r.Generate(out); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p.(*parse.Count).Last(), "records are processed.")
}
