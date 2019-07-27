// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// always put all the related things together as in here

// result stores metrics for a domain
// it uses the value mechanics,
// because it doesn't have to update anything
type result struct {
	domain string
	visits int
	// add more metrics if needed
}

// add adds the metrics of another result to itself and returns a new Result
func addResult(r result, other result) result {
	return result{
		domain: r.domain,
		visits: r.visits + other.visits,
	}
}
