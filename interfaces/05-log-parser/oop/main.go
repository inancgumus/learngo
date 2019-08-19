// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"log"
	// "fmt"
	"os"
)

func main() {
	a := newAnalysis()
	// a.filterBy(notUsing(domainExtFilter("io", "com")))
	// a.groupBy(domainGrouper)

	p := newTextParser(os.Stdin)
	s := newTextSummary()

	// s := &chartSummary{
  	//  title:  "visits per domain",
  	//  width:  1920,
  	//  height: 800,
  	// }

	if err := report(p, a, s); err != nil {
		log.Fatalln(err)
	}

	// if err := reportFromFile(os.Args[1]); err != nil {
	// 	log.Fatalln(err)
	// }
}
