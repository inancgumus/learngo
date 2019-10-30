// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package report

import (
	"sort"
)

// Summary aggregates the parsing results
type Summary struct {
	sum     map[string]Result // metrics per domain
	domains []string          // unique domain names
	total   Result            // total visits for all domains
}

// newSummary constructs and initializes a new summary
// You can't use its methods without pointer mechanics
func newSummary() *Summary {
	return &Summary{sum: make(map[string]Result)}
}

// Update updates the report for the given parsing result
func (s *Summary) update(r Result) {
	domain := r.Domain
	if _, ok := s.sum[domain]; !ok {
		s.domains = append(s.domains, domain)
	}

	// let the result handle the addition
	// this allows us to manage the result in once place
	// and this way it becomes easily extendable
	s.total = s.total.add(r)
	s.sum[domain] = r.add(s.sum[domain])
}

// Iterator returns `next()` to detect when the iteration ends,
// and a `cur()` to return the current result.
// iterator iterates sorted by domains.
func (s *Summary) Iterator() (next func() bool, cur func() Result) {
	sort.Strings(s.domains)

	// remember the last iterated result
	var last int

	next = func() bool {
		defer func() { last++ }()
		return len(s.domains) > last
	}

	cur = func() Result {
		// returns a copy so the caller cannot change it
		name := s.domains[last-1]
		return s.sum[name]
	}

	return
}

// Total returns the total metrics
func (s *Summary) Total() Result {
	return s.total
}

// For the interfaces section
//
// MarshalJSON marshals a report to JSON
// Alternative: unexported embedding
// func (s *Summary) MarshalJSON() ([]byte, error) {
// 	type total struct {
// 		*Result
// 		IgnoreDomain *string `json:"domain,omitempty"`
// 	}

// 	return json.Marshal(struct {
// 		Sum     map[string]Result `json:"summary"`
// 		Domains []string          `json:"domains"`
// 		Total   total             `json:"total"`
// 	}{
// 		Sum: s.sum, Domains: s.domains, Total: total{Result: &s.total},
// 	})
// }
