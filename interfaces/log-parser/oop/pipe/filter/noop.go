// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package filter

import "github.com/inancgumus/learngo/interfaces/log-parser/oop/pipe"

// Noop filter that does nothing.
func Noop(r pipe.Record) bool {
	return true
}
