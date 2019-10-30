// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

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

// Consume the records for lazy filtering.
func (f *Filter) Consume(records Iterator) error {
	f.src = records
	return nil
}

// Each filtered records.
func (f *Filter) Each(yield func(Record) error) error {
	records := func(r Record) error {
		if !f.checkAll(r) {
			return nil
		}
		return yield(r)
	}

	return f.src.Each(records)
}

// checkAll the filters against the record.
func (f *Filter) checkAll(r Record) bool {
	for _, fi := range f.filters {
		if !fi(r) {
			return false
		}
	}
	return true
}
