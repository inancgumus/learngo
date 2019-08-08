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
