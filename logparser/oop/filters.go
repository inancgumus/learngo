package main

import "strings"

type filterFunc func(record) bool

func noopFilter(r record) bool {
	return true
}

func notUsing(filter filterFunc) filterFunc {
	return func(r record) bool {
		return !filter(r)
	}
}

func domainExtFilter(domains ...string) filterFunc {
	return func(r record) bool {
		for _, domain := range domains {
			if strings.HasSuffix(r.domain, "."+domain) {
				return true
			}
		}
		return false
	}
}

func domainFilter(domain string) filterFunc {
	return func(r record) bool {
		return strings.Contains(r.domain, domain)
	}
}

func orgDomainsFilter(r record) bool {
	return strings.HasSuffix(r.domain, ".org")
}
