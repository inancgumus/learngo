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

// transform the record
func (f *filter) digest(records iterator) error {
	f.src = records
	return nil
}

// each yields only the filtered records
func (f *filter) each(yield recordFn) error {
	return f.src.each(func(r record) {
		if !f.check(r) {
			return
		}
		yield(r)
	})
}

// check all the filters against the record
func (f *filter) check(r record) bool {
	for _, fi := range f.filters {
		if !fi(r) {
			return false
		}
	}
	return true
}
