// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// logCount counts the yielded records
type logCount struct {
	iterator
	n int
}

func (lc *logCount) each(yield recordFn) error {
	err := lc.iterator.each(func(r record) {
		lc.n++
		yield(r)
	})

	if err != nil {
		// lc.n+1: iterator.each won't call yield on err
		return fmt.Errorf("record %d: %v", lc.n+1, err)
	}
	return nil
}

func (lc *logCount) count() int {
	return lc.n
}
