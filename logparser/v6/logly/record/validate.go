// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package record

import "errors"

// validate whether the current record is valid or not.
func (r *Record) validate() error {
	var msg string

	switch {
	case r.Domain == "":
		msg = "record.domain cannot be empty"
	case r.Page == "":
		msg = "record.page cannot be empty"
	case r.Visits < 0:
		msg = "record.visits cannot be negative"
	case r.Uniques < 0:
		msg = "record.uniques cannot be negative"
	}

	if msg != "" {
		return errors.New(msg)
	}
	return nil
}
