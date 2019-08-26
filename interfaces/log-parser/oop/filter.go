// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

type filter struct {
	src     iterator
	filters []filterFunc
}

func filterBy(fn ...filterFunc) *filter {
	return &filter{filters: fn}
}

// transform the result
func (f *filter) digest(results iterator) error {
	f.src = results
	return nil
}

// each yields an analysis result
func (f *filter) each(yield resultFn) error {
	return f.src.each(func(r result) {
		if !f.check(r) {
			return
		}
		yield(r)
	})
}

// check all the filters against the result
func (f *filter) check(r result) bool {
	for _, fi := range f.filters {
		if !fi(r) {
			return false
		}
	}
	return true
}
