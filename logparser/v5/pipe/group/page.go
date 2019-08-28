// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package group

import "github.com/inancgumus/learngo/logparser/v5/pipe"

// Page groups records by page.
func Page(r pipe.Record) string {
	return r.Str("domain") + r.Str("page")
}
