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
	an := newAnalysis()
	an.filterBy(notUsing(domainExtFilter("io", "com")))
	an.groupBy(domainGrouper)

	// pars, err := parseTextFile("log.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := report(an, newTextParser(os.Stdin), summarizeFunc(textSummary))
	// err := report(an, newJSONParser(os.Stdin), newTextSummary())

	// chart := &chartSummary{
	// 	title:  "visits per domain",
	// 	width:  1920,
	// 	height: 800,
	// }

	// err := report(an, newTextParser(os.Stdin), chart)

	if err != nil {
		log.Fatalln(err)
	}
}
