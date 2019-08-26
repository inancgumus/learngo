package main

import (
	"encoding/json"
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

func (r result) add(other result) result {
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

	r.visits, err = strconv.Atoi(fields[2])
	if err != nil || r.visits < 0 {
		return fmt.Errorf("wrong input %q", fields[2])
	}

	r.uniques, err = strconv.Atoi(fields[3])
	if err != nil || r.uniques < 0 {
		return fmt.Errorf("wrong input %q", fields[3])
	}
	return nil
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

	return nil
}
