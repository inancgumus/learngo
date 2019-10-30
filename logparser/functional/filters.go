// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "strings"

func noopFilter(r result) bool {
	return true
}

func notUsing(filter filterFn) filterFn {
	return func(r result) bool {
		return !filter(r)
	}
}

func domainExtFilter(domains ...string) filterFn {
	return func(r result) bool {
		for _, domain := range domains {
			if strings.HasSuffix(r.domain, "."+domain) {
				return true
			}
		}
		return false
	}
}

func domainFilter(domain string) filterFn {
	return func(r result) bool {
		return strings.Contains(r.domain, domain)
	}
}

func orgDomainsFilter(r result) bool {
	return strings.HasSuffix(r.domain, ".org")
}
