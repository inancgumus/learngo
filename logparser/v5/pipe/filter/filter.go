// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package filter

import "github.com/inancgumus/learngo/logparser/v5/pipe"

// Func represents a filtering pipeline func.
type Func func(pipe.Record) (pass bool)

// Filter the records.
type Filter struct {
	src     pipe.Iterator
	filters []Func
}

// By returns a new filter pipeline.
func By(fn ...Func) *Filter {
	return &Filter{filters: fn}
}

// Consume saves the iterator for later processing.
func (f *Filter) Consume(records pipe.Iterator) error {
	f.src = records
	return nil
}

// Each yields only the filtered records.
func (f *Filter) Each(yield func(pipe.Record)) error {
	return f.src.Each(func(r pipe.Record) {
		if !f.check(r) {
			return
		}
		yield(r)
	})
}

// check all the filters against the record.
func (f *Filter) check(r pipe.Record) bool {
	for _, fi := range f.filters {
		if !fi(r) {
			return false
		}
	}
	return true
}
