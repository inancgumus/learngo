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
	r      io.Reader
	update func(r result)
}

func newJSONParser(r io.Reader) *jsonParser {
	return &jsonParser{
		r:      r,
		update: func(result) {},
	}
}

func (p *jsonParser) parse() error {
	defer readClose(p.r)

	bytes, err := ioutil.ReadAll(p.r)
	if err != nil {
		return err
	}

	return p.parseJSON(bytes)
}

func (p *jsonParser) parseJSON(bytes []byte) error {
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
		p.update(result{
			domain:  r.Domain,
			page:    r.Page,
			visits:  r.Visits,
			uniques: r.Uniques,
		})
	}

	return nil
}

func (p *jsonParser) notify(f func(r result)) {
	p.update = f
}
