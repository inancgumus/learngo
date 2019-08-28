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
)

func main() {
	// newGrouper(domainGrouper)

	// s := &chartReport{
	//  title:  "visits per domain",
	//  width:  1920,
	//  height: 800,
	// }

	// pipe, err := fromFile("../logs/log.jsonl")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	pipe := newPipeline(
		newTextLog(os.Stdin),
		newTextReport(),
		filterBy(notUsing(domainExtFilter("com", "io"))),
		groupBy(domainGrouper),
	)

	if err := pipe.run(); err != nil {
		log.Fatalln(err)
	}

	// if err := reportFromFile(os.Args[1]); err != nil {
	// 	log.Fatalln(err)
	// }
}
