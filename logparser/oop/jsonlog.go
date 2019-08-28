// For more tutorials: https://bj.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"encoding/json"
	"io"
)

type jsonLog struct {
	reader io.Reader
}

func newJSONLog(r io.Reader) *jsonLog {
	return &jsonLog{reader: r}
}

func (j *jsonLog) each(yield recordFn) error {
	defer readClose(j.reader)

	dec := json.NewDecoder(bufio.NewReader(j.reader))

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
