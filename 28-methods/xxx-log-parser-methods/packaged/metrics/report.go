// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package metrics

import (
	"encoding/json"
	"sort"
)

// Report aggregates the final report
type Report struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   result            // total visits for all domains
}

// NewReport constructs and initializes a new report
// You can't use its methods without pointer mechanics
func NewReport() *Report {
	return &Report{sum: make(map[string]result)}
}

// Update updates the report for the given parsing result
func (r *Report) Update(p Parsed) {
	// do not update the report if the result has an error
	if p.err != nil {
		return
	}

	domain := p.Domain
	if _, ok := r.sum[domain]; !ok {
		r.domains = append(r.domains, domain)
	}

	// let the result handle the addition
	// this allows us to manage the result in once place
	// and this way it becomes easily extendable
	r.total = r.total.add(p.result)
	r.sum[domain] = p.add(r.sum[domain])
}

// Iterator returns `next()` to detect when the iteration ends,
// and a `cur()` to return the current result.
// iterator iterates sorted by domains.
func (r *Report) Iterator() (next func() bool, cur func() result) {
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

// Total returns the total metrics
func (r *Report) Total() Parsed {
	return Parsed{result: r.total}
}

// MarshalJSON marshals a report to JSON
// Alternative: unexported embedding
func (r *Report) MarshalJSON() ([]byte, error) {
	type total struct {
		*result
		IgnoreDomain *string `json:"Domain,omitempty"`
	}

	return json.Marshal(struct {
		Sum     map[string]result
		Domains []string
		Total   total
	}{
		Sum: r.sum, Domains: r.domains, Total: total{result: &r.total},
	})
}
