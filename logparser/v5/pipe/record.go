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
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const fieldsLength = 4

// record stores fields of a log line.
type record struct {
	domain  string
	page    string
	visits  int
	uniques int
}

// recordJSON is used for marshaling and unmarshaling JSON.
type recordJSON struct {
	Domain  string
	Page    string
	Visits  int
	Uniques int
}

// sum the numeric fields with another record.
func (r record) sum(other record) record {
	r.visits += other.visits
	r.uniques += other.uniques
	return r
}

// UnmarshalText to a *record.
func (r *record) UnmarshalText(p []byte) (err error) {
	fields := strings.Fields(string(p))

	if len(fields) != fieldsLength {
		return fmt.Errorf("wrong number of fields %q", fields)
	}

	r.domain, r.page = fields[0], fields[1]

	if r.visits, err = parseStr("visits", fields[2]); err != nil {
		return err
	}
	if r.uniques, err = parseStr("uniques", fields[3]); err != nil {
		return err
	}

	return validate(*r)
}

// UnmarshalJSON to a *record.
func (r *record) UnmarshalJSON(data []byte) error {
	var rj recordJSON

	if err := json.Unmarshal(data, &rj); err != nil {
		return err
	}

	*r = record{rj.Domain, rj.Page, rj.Visits, rj.Uniques}

	return validate(*r)
}

// MarshalJSON of a *record.
func (r *record) MarshalJSON() ([]byte, error) {
	rj := recordJSON{r.domain, r.page, r.visits, r.uniques}
	return json.Marshal(rj)
}

// parseStr helps UnmarshalText for string to positive int parsing.
func parseStr(name, v string) (int, error) {
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("Record.UnmarshalText %q: %v", name, err)
	}
	return n, nil
}

// validate whether a parsed record is valid or not.
func validate(r record) (err error) {
	switch {
	case r.domain == "":
		err = errors.New("record.domain cannot be empty")
	case r.page == "":
		err = errors.New("record.page cannot be empty")
	case r.visits < 0:
		err = errors.New("record.visits cannot be negative")
	case r.uniques < 0:
		err = errors.New("record.uniques cannot be negative")
	}
	return
}
