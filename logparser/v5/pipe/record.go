// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pipe

// Record provides a generic interface for any sort of records.
type Record interface {
	Str(field string) string
	Int(field string) int
}

// Summer provides a method for summing the numeric fields.
type Summer interface {
	Sum(Record) Record
}
