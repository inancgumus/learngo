// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"sort"
	"strings"
)

// product is now *product because sorting will unnecessarily copy the product values

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚.\n"
	}

	sort.Sort(l)
	var str strings.Builder
	for _, p := range l {
		fmt.Fprintf(&str, "* %s\n", p)
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

// by default `list` sorts by `title`.
func (l list) Len() int           { return len(l) }
func (l list) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l list) Less(i, j int) bool { return l[i].title < l[j].title }

// byRelease sorts by product release dates.
type byRelease struct {
	// byRelease embeds `list` and reuses list's Len and Swap methods.
	list
}

func (bp byRelease) Less(i, j int) bool {
	// Before() accepts a time.Time but `released` is not time.Time.
	// `released.Time` is necessary.
	return bp.list[i].released.Before(bp.list[j].released.Time)
}
