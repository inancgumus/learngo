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
	sum  map[string]result // metrics per group key
	keys []string          // unique group keys
	key  groupFunc
}

func groupBy(key groupFunc) *group {
	return &group{
		sum: make(map[string]result),
		key: key,
	}
}

// digest the results
func (g *group) digest(results iterator) error {
	return results.each(func(r result) {
		k := g.key(r)

		if _, ok := g.sum[k]; !ok {
			g.keys = append(g.keys, k)
		}

		g.sum[k] = r.add(g.sum[k])
	})
}

// each yields the grouped results
func (g *group) each(yield resultFn) error {
	sort.Strings(g.keys)

	for _, k := range g.keys {
		yield(g.sum[k])
	}

	return nil
}
