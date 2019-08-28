// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pipe

import "fmt"

// logCount counts the yielded records.
type logCount struct {
	Iterator
	n int
}

// Each yields to the inner iterator while counting the records.
// Reports the record number on an error.
func (lc *logCount) Each(yield func(Record)) error {
	err := lc.Iterator.Each(func(r Record) {
		lc.n++
		yield(r)
	})

	if err != nil {
		// lc.n+1: iterator.each won't call yield on err
		return fmt.Errorf("record %d: %v", lc.n+1, err)
	}
	return nil
}

// count returns the last read record number.
func (lc *logCount) count() int {
	return lc.n
}
