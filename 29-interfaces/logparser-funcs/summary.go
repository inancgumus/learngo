// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"sort"
)

// summary aggregates the parsing results
type summary struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   result            // total visits for all domains
}

// newSummary constructs and initializes a new summary
// You can't use its methods without pointer mechanics
func newSummary() *summary {
	return &summary{sum: make(map[string]result)}
}

// updateSummary updates the report for the given parsing result
func updateSummary(s *summary, r result) {
	domain := r.domain
	if _, ok := s.sum[domain]; !ok {
		s.domains = append(s.domains, domain)
	}

	// let the result handle the addition
	// this allows us to manage the result in once place
	// and this way it becomes easily extendable
	s.total = addResult(s.total, r)
	s.sum[domain] = addResult(r, s.sum[domain])
}

// iteratorSummary returns `next()` to detect when the iteration ends,
// and a `cur()` to return the current result.
// iterator iterates sorted by domains.
func iteratorSummary(s *summary) (next func() bool, cur func() result) {
	sort.Strings(s.domains)

	// remember the last iterated result
	var last int

	next = func() bool {
		// done := len(s.domains) > last
		// last++
		// return done
		defer func() { last++ }()
		return len(s.domains) > last
	}

	cur = func() result {
		// returns a copy so the caller cannot change it
		name := s.domains[last-1]
		return s.sum[name]
	}

	return
}

// totalsSummary returns the total metrics
func totalsSummary(s *summary) result {
	return s.total
}
