// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "sort"

type (
	groupFunc  func(result) string
	filterFunc func(result) bool
)

type analysis struct {
	sum      map[string]result // metrics per domain
	keys     []string          // unique keys
	groupKey groupFunc
	filter   filterFunc
}

func newAnalysis() *analysis {
	return &analysis{
		sum:      make(map[string]result),
		groupKey: domainGrouper,
		filter:   noopFilter,
	}
}

func (a *analysis) groupBy(g groupFunc)   { a.groupKey = g }
func (a *analysis) filterBy(f filterFunc) { a.filter = f }

func (a *analysis) analyse(r result) {
	if !a.filter(r) {
		return
	}

	key := a.groupKey(r)
	if _, ok := a.sum[key]; !ok {
		a.keys = append(a.keys, key)
	}

	a.sum[key] = r.add(a.sum[key])
}

// each analysed result will be sent by the given func
func (a *analysis) each(f func(r result)) {
	sort.Strings(a.keys)

	for _, domain := range a.keys {
		f(a.sum[domain])
	}
}
