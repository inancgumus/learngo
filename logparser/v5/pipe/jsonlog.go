// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package pipe

import (
	"encoding/json"
	"io"
)

// JSON parses json records.
type JSON struct {
	reader io.Reader
}

// NewJSONLog creates a json parser.
func NewJSONLog(r io.Reader) *JSON {
	return &JSON{reader: r}
}

// Each sends the records from a reader to upstream.
func (j *JSON) Each(yield func(Record) error) error {
	defer readClose(j.reader)

	// Use the same record for unmarshaling.
	var r Record

	dec := json.NewDecoder(j.reader)

	for {
		err := dec.Decode(&r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := yield(r); err != nil {
			return err
		}
	}

	return nil
}
