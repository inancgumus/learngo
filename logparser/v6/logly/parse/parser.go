// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package parse

import "github.com/inancgumus/learngo/logparser/v6/logly/record"

// Parser is an interface for the parsers.
type Parser interface {
	// Parse the next record from the source.
	Parse() bool

	// Value returns the last parsed record by a call to Parse.
	Value() record.Record

	// Err returns the first error that was encountered.
	Err() error
}
