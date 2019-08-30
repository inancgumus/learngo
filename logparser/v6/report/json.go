// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package report

import (
	"encoding/json"
	"io"

	"github.com/inancgumus/learngo/logparser/v6/parse"
)

// JSONReport generates a JSON report.
type JSONReport struct {
	w *json.Encoder
}

// JSON returns a JSON report generator.
func JSON(w io.Writer) *JSONReport {
	return &JSONReport{
		w: json.NewEncoder(w),
	}
}

// Generate the report from the records.
func (jr *JSONReport) Generate(rs []parse.Record) error {
	for _, r := range rs {
		if err := jr.w.Encode(&r); err != nil {
			return err
		}
	}
	return nil
}
