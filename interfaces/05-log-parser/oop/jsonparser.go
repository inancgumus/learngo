// For more tutorials: https://bp.learngoprogramming.com
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

type jsonParser struct {
	r io.Reader
}

func newJSONParser(r io.Reader) *jsonParser {
	return &jsonParser{r: r}
}

func (p *jsonParser) parse(handle resultFn) error {
	defer readClose(p.r)

	bytes, err := ioutil.ReadAll(p.r)
	if err != nil {
		return err
	}

	return parseJSON(bytes, handle)
}

func parseJSON(bytes []byte, handle resultFn) error {
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
		handle(result{
			domain:  r.Domain,
			page:    r.Page,
			visits:  r.Visits,
			uniques: r.Uniques,
		})
	}

	return nil
}
