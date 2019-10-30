// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package record

// Sum groups the records by summing their numeric fields.
type Sum struct {
	sum map[string]Record
}

// SumGroup the records by domain.
func SumGroup() *Sum {
	return &Sum{
		sum: make(map[string]Record),
	}
}

// Group the record.
func (s *Sum) Group(r Record) {
	k := r.Domain
	r.Sum(s.sum[k])
	s.sum[k] = r
}

// Records returns the grouped records.
func (s *Sum) Records() []Record {
	var out []Record
	for _, res := range s.sum {
		out = append(out, res)
	}
	return out
}
