package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const fieldsLength = 4

type result struct {
	domain  string
	page    string
	visits  int
	uniques int
}

func (r result) sum(other result) result {
	r.visits += other.visits
	r.uniques += other.uniques
	return r
}

// UnmarshalText to a *result
func (r *result) UnmarshalText(p []byte) (err error) {
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

// UnmarshalJSON to a *result
func (r *result) UnmarshalJSON(data []byte) error {
	var re struct {
		Domain  string
		Page    string
		Visits  int
		Uniques int
	}

	if err := json.Unmarshal(data, &re); err != nil {
		return err
	}

	*r = result{re.Domain, re.Page, re.Visits, re.Uniques}
	return validate(*r)
}

// parseStr helps UnmarshalText for string to positive int parsing
func parseStr(name, v string) (int, error) {
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("result.UnmarshalText %q: %v", name, err)
	}
	return n, nil
}

func validate(r result) (err error) {
	switch {
	case r.domain == "":
		err = errors.New("result.domain cannot be empty")
	case r.page == "":
		err = errors.New("result.page cannot be empty")
	case r.visits < 0:
		err = errors.New("result.visits cannot be negative")
	case r.uniques < 0:
		err = errors.New("result.uniques cannot be negative")
	}
	return
}
