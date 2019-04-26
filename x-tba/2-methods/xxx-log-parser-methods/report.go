// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

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
func (r *report) update(parsed parserResult) {
	// do not update the report if the result has an error
	if parsed.err != nil {
		return
	}

	domain := parsed.domain
	if _, ok := r.sum[domain]; !ok {
		r.domains = append(r.domains, domain)
	}

	// let the result handle the addition
	// this allows us to manage the result in once place
	// and this way it becomes easily extendable
	r.total = r.total.add(parsed.result)
	r.sum[domain] = parsed.add(r.sum[domain])
}
