// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package group

import "github.com/inancgumus/learngo/logparser/v6/pipe"

// Domain groups the records by domain.
// It keeps the other fields intact.
// For example: It returns the page field as well.
// Exercise: Write a solution that removes the unnecessary data.
func Domain(r pipe.Record) string {
	return r.Str("domain")
}
