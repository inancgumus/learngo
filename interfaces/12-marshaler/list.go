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

// by default `list` sorts by `Title`.
func (l list) Len() int           { return len(l) }
func (l list) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l list) Less(i, j int) bool { return l[i].Title < l[j].Title }

// byRelease sorts by product release dates.
type byRelease struct {
	list
}

func (bp byRelease) Less(i, j int) bool {
	return bp.list[i].Released.Before(bp.list[j].Released.Time)
}
