// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package filter

import (
	"strings"

	"github.com/inancgumus/learngo/logparser/v5/pipe"
)

// DomainExt filters a set of domain extensions.
func DomainExt(domains ...string) Func {
	return func(r pipe.Record) bool {
		for _, domain := range domains {
			if strings.HasSuffix(r.Str("domain"), "."+domain) {
				return true
			}
		}
		return false
	}
}

// Domain filters a domain if it contains the given text.
func Domain(text string) Func {
	return func(r pipe.Record) bool {
		return strings.Contains(r.Str("domain"), text)
	}
}

// DomainOrg filters only the ".org" domains.
func DomainOrg(r pipe.Record) bool {
	return strings.HasSuffix(r.Str("domain"), ".org")
}
