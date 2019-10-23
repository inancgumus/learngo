// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"sort"
	"strings"
)

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚.\n"
	}

	var str strings.Builder
	for _, p := range l {
		// fmt.Printf("* %s\n", p)
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

// implementation of the sort.Interface:

// by default `list` sorts by `Title`.
func (l list) Len() int           { return len(l) }
func (l list) Less(i, j int) bool { return l[i].Title < l[j].Title }
func (l list) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

// byRelease sorts by product release dates.
type byRelease struct{ list }

func (bp byRelease) Less(i, j int) bool {
	return bp.list[i].Released.Before(bp.list[j].Released.Time)
}

func byReleaseDate(l list) sort.Interface {
	return byRelease{l}
}
