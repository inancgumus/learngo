// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "sort"

// report aggregates the final report
type report struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   result            // total visits for all domains
}

// newReport constructs and initializes a new report
func newReport() *report {
	return &report{sum: make(map[string]result)}
}

// update updates the errors for the given parsing result
func (r *report) update(p parsed) {
	// do not update the report if the result has an error
	if p.err != nil {
		return
	}

	domain := p.domain
	if _, ok := r.sum[domain]; !ok {
		r.domains = append(r.domains, domain)
	}

	// let the result handle the addition
	// this allows us to manage the result in once place
	// and this way it becomes easily extendable
	r.total = r.total.add(p.result)
	r.sum[domain] = p.add(r.sum[domain])
}

// iterator returns `next()` to detect when the iteration ends,
// and a `cur()` to return the current result.
// iterator iterates sorted by domains.
func (r *report) iterator() (next func() bool, cur func() result) {
	sort.Strings(r.domains)

	// remember the last iterated result
	var last int

	next = func() bool {
		defer func() { last++ }()
		return len(r.domains) > last
	}

	cur = func() result {
		// returns a copy so the caller cannot change it
		name := r.domains[last-1]
		return r.sum[name]
	}

	return
}
