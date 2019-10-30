// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package report

import (
	"encoding/json"
	"io"

	"github.com/inancgumus/learngo/logparser/v6/logly/record"
)

// JSON generates a json report.
func JSON(w io.Writer, rs []record.Record) error {
	enc := json.NewEncoder(w)

	for _, r := range rs {
		err := enc.Encode(&r)

		if err != nil {
			return err
		}
	}
	return nil
}
