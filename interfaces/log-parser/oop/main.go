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
	an := newAnalysis()
	// an.filterBy(notUsing(domainExtFilter("io", "com")))
	// an.filterBy(domainFilter("org"))
	// an.groupBy(domainGrouper)

	src := newTextLog(os.Stdin)
	dst := newTextReport()

	// s := &chartReport{
	//  title:  "visits per domain",
	//  width:  1920,
	//  height: 800,
	// }

	pipe := newPipeline(src, dst, an)

	if err := pipe.run(); err != nil {
		log.Fatalln(err)
	}

	// if err := reportFromFile(os.Args[1]); err != nil {
	// 	log.Fatalln(err)
	// }
}
