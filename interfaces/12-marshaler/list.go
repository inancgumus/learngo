// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

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

func (br byRelease) Less(i, j int) bool {
	return br.list[i].Released.Before(br.list[j].Released.Time)
}

func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}
