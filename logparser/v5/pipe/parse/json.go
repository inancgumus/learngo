// For more tutorials: https://bj.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package parse

import (
	"encoding/json"
	"io"

	"github.com/inancgumus/learngo/logparser/v5/pipe"
)

// JSON parses json records.
type JSON struct {
	reader io.Reader
}

// FromJSON creates a json parser.
func FromJSON(r io.Reader) *JSON {
	return &JSON{reader: r}
}

// Each yields records from a json reader.
func (j *JSON) Each(yield func(pipe.Record)) error {
	defer readClose(j.reader)

	dec := json.NewDecoder(j.reader)

	for {
		var r record

		err := dec.Decode(&r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		yield(r)
	}
	return nil
}
