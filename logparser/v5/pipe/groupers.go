// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package pipe

// DomainGrouper groups the records by domain.
// It keeps the other fields intact.
// For example: It returns the page field as well.
// Exercise: Write a solution that removes the unnecessary data.
func DomainGrouper(r Record) string {
	return r.domain
}

// Page groups records by page.
func Page(r Record) string {
	return r.domain + r.page
}
