// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pipe

// FilterFunc represents a filtering pipeline func.
// The type alias frees us from binding to a named type.
type FilterFunc = func(Record) (pass bool)

// Filter the records.
type Filter struct {
	src     Iterator
	filters []FilterFunc
}

// FilterBy returns a new filter pipeline.
func FilterBy(fn ...FilterFunc) *Filter {
	return &Filter{filters: fn}
}

// Consume saves the iterator for later processing.
func (f *Filter) Consume(records Iterator) error {
	f.src = records
	return nil
}

// Each yields only the filtered records.
func (f *Filter) Each(yield func(Record) error) error {
	return f.src.Each(func(r Record) error {
		if !f.check(r) {
			return nil
		}
		return yield(r)
	})
}

// check all the filters against the record.
func (f *Filter) check(r Record) bool {
	for _, fi := range f.filters {
		if !fi(r) {
			return false
		}
	}
	return true
}
