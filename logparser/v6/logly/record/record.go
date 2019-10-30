// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package record

const fieldsLength = 4

// Record stores fields of a log line.
type Record struct {
	Domain  string
	Page    string
	Visits  int
	Uniques int
}

// Sum the numeric fields with another record.
func (r *Record) Sum(other Record) {
	r.Visits += other.Visits
	r.Uniques += other.Uniques
}

// Reset all the fields of this record.
func (r *Record) Reset() {
	r.Domain = ""
	r.Page = ""
	r.Visits = 0
	r.Uniques = 0
}
