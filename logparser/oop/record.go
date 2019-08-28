package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const fieldsLength = 4

type record struct {
	domain  string
	page    string
	visits  int
	uniques int
}

func (r record) sum(other record) record {
	r.visits += other.visits
	r.uniques += other.uniques
	return r
}

// UnmarshalText to a *record
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

// UnmarshalJSON to a *record
func (r *record) UnmarshalJSON(data []byte) error {
	var re struct {
		Domain  string
		Page    string
		Visits  int
		Uniques int
	}

	if err := json.Unmarshal(data, &re); err != nil {
		return err
	}

	*r = record{re.Domain, re.Page, re.Visits, re.Uniques}
	return validate(*r)
}

// parseStr helps UnmarshalText for string to positive int parsing
func parseStr(name, v string) (int, error) {
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("record.UnmarshalText %q: %v", name, err)
	}
	return n, nil
}

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
