// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "strings"

func filter(r result) bool {
	return filterOrg(r)
}

func filterOrg(r result) bool {
	return strings.HasSuffix(r.domain, ".org")
}

func filterBlogs(r result) bool {
	return strings.HasPrefix(r.domain, "blog.")
}

// type filterFunc func(result) bool

// func filter(r result, process filterFunc) bool {
// 	return process(r)
// }

// func noopFilter(r result) bool {
// 	return true
// }

// func notUsing(filter filterFunc) filterFunc {
// 	return func(r result) bool {
// 		return !filter(r)
// 	}
// }
