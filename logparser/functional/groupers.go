// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// domainGrouper groups by domain.
// but it keeps the other fields.
// for example: it returns pages as well, but you shouldn't use them.
// exercise: write a function that erases superfluous data.
func domainGrouper(r result) string {
	return r.domain
}

func pageGrouper(r result) string {
	return r.domain + r.page
}

// groupBy allocates map unnecessarily
func noopGrouper(r result) string {
	// with something like:
	// return randomStrings()
	return ""
}
