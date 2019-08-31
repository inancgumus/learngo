// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package record

import "encoding/json"

// UnmarshalJSON to a record.
func (r *Record) UnmarshalJSON(data []byte) error {
	type rjson Record

	err := json.Unmarshal(data, (*rjson)(r))
	if err != nil {
		return err
	}

	return r.validate()
}
