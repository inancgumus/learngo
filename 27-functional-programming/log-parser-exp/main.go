// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	defer recoverErr()

	_, err := newPipeline().
		// from(fastTextReader(os.Stdin)).
		filterBy(notUsing(domainExtFilter("com", "io"))).
		groupBy(domainGrouper).
		start()

	if err != nil {
		fmt.Println("> Err:", err)
	}
}

func recoverErr() {
	val := recover()

	if val == nil {
		return
	}

	fmt.Println("> Error occurred:", val)
}

/*
newReport -> report.New().
Result    -> report.Line

notUsing = report.Not

pl := newPipeline(pipeOpts{
  from:     fastTextReader(os.Stdin),
  filterBy: notUsing(domainExtFilter("com", "io")),
  groupBy:  domainGrouper,
})

err := pl.start()

_, err := report.New().
	From(report.TextReader(os.Stdin)).
	To(report.TextWriter(os.Stdout)).
	// FilterBy(report.OrgDomainsFilter).
	FilterBy(notUsing(report.DomainExtFilter("com", "io"))).
	GroupBy(report.DomainGrouper).
	Start()
*/
