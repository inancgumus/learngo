// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package filter

import "github.com/inancgumus/learngo/logparser/v6/pipe"

// Not reverses a filter. True becomes false, and vice versa.
func Not(filter Func) Func {
	return func(r pipe.Record) bool {
		return !filter(r)
	}
}
