// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pipe

import "strings"

// NotFilter reverses a filter. True becomes false, and vice versa.
func NotFilter(filter FilterFunc) FilterFunc {
	return func(r Record) bool {
		return !filter(r)
	}
}

// DomainExtFilter filters a set of domain extensions.
func DomainExtFilter(domains ...string) FilterFunc {
	return func(r Record) bool {
		for _, domain := range domains {
			if strings.HasSuffix(r.domain, "."+domain) {
				return true
			}
		}
		return false
	}
}

// DomainFilter filters a domain if it contains the given text.
func DomainFilter(text string) FilterFunc {
	return func(r Record) bool {
		return strings.Contains(r.domain, text)
	}
}

// DomainOrgFilter filters only the ".org" domains.
func DomainOrgFilter(r Record) bool {
	return strings.HasSuffix(r.domain, ".org")
}
