// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"sort"
)

type group struct {
	sum  map[string]record // metrics per group key
	keys []string          // unique group keys
	key  groupFunc
}

func groupBy(key groupFunc) *group {
	return &group{
		sum: make(map[string]record),
		key: key,
	}
}

// digest the records
func (g *group) digest(records iterator) error {
	return records.each(func(r record) {
		k := g.key(r)

		if _, ok := g.sum[k]; !ok {
			g.keys = append(g.keys, k)
		}

		g.sum[k] = r.sum(g.sum[k])
	})
}

// each yields the grouped records
func (g *group) each(yield recordFn) error {
	sort.Strings(g.keys)

	for _, k := range g.keys {
		yield(g.sum[k])
	}

	return nil
}
