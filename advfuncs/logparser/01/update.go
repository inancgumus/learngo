// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// update updates all the parsing results using the given parsing result
func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}

	if !filter(r) {
		return
	}

	// Collect the unique domains
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	// Keep track of total and per domain visits
	p.total += r.visits

	// create and assign a new copy of `visit`
	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}
}
