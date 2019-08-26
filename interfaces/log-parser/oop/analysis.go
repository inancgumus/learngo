// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "sort"

type analysis struct {
	sum      map[string]result // metrics per group key
	keys     []string          // unique group keys
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

// transform the result
func (a *analysis) transform(r result) {
	if !a.filter(r) {
		return
	}

	key := a.groupKey(r)
	if _, ok := a.sum[key]; !ok {
		a.keys = append(a.keys, key)
	}

	a.sum[key] = r.add(a.sum[key])
}

// each yields an analysis result
func (a *analysis) each(yield resultFn) error {
	sort.Strings(a.keys)

	for _, key := range a.keys {
		yield(a.sum[key])
	}

	return nil
}

func (a *analysis) groupBy(g groupFunc) {
	if g != nil {
		a.groupKey = g
	}
}

func (a *analysis) filterBy(f filterFunc) {
	if f != nil {
		a.filter = f
	}
}
