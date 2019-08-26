// For more tutorials: https://bj.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type jsonLog struct {
	reader io.Reader
}

func newJSONLog(r io.Reader) *jsonLog {
	return &jsonLog{reader: r}
}

func (j *jsonLog) each(yield resultFn) error {
	defer readClose(j.reader)

	bytes, err := ioutil.ReadAll(j.reader)
	if err != nil {
		return err
	}

	return extractJSON(bytes, yield)
}

func extractJSON(bytes []byte, yield resultFn) error {
	var rs []struct {
		Domain  string
		Page    string
		Visits  int
		Uniques int
	}

	if err := json.Unmarshal(bytes, &rs); err != nil {
		if serr, ok := err.(*json.SyntaxError); ok {
			return fmt.Errorf("%v %q", serr, bytes[:serr.Offset])
		}
		return err
	}

	for _, r := range rs {
		yield(result{
			domain:  r.Domain,
			page:    r.Page,
			visits:  r.Visits,
			uniques: r.Uniques,
		})
	}

	return nil
}
