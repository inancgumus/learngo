package parse

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const fieldsLength = 4

// Record stores fields of a log line.
type Record struct {
	Domain  string
	Page    string
	Visits  int
	Uniques int
}

// Sum the numeric fields with another record.
func (r *Record) Sum(other Record) {
	r.Visits += other.Visits
	r.Uniques += other.Uniques
}

// Reset all the fields of this record.
func (r *Record) Reset() {
	r.Domain = ""
	r.Page = ""
	r.Visits = 0
	r.Uniques = 0
}

// FromText unmarshals the log line into this record.
func (r *Record) FromText(p []byte) (err error) {
	fields := strings.Fields(string(p))

	if len(fields) != fieldsLength {
		return fmt.Errorf("wrong number of fields %q", fields)
	}

	r.Domain = fields[0]
	r.Page = fields[1]

	const msg = "record.UnmarshalText %q: %v"
	if r.Visits, err = strconv.Atoi(fields[2]); err != nil {
		return fmt.Errorf(msg, "visits", err)
	}
	if r.Uniques, err = strconv.Atoi(fields[3]); err != nil {
		return fmt.Errorf(msg, "uniques", err)
	}

	return r.validate()
}

// UnmarshalJSON to a record.
func (r *Record) UnmarshalJSON(data []byte) error {
	type rjson Record

	err := json.Unmarshal(data, (*rjson)(r))
	if err != nil {
		return err
	}

	return r.validate()
}

// validate whether the current record is valid or not.
func (r *Record) validate() error {
	var msg string

	switch {
	case r.Domain == "":
		msg = "record.domain cannot be empty"
	case r.Page == "":
		msg = "record.page cannot be empty"
	case r.Visits < 0:
		msg = "record.visits cannot be negative"
	case r.Uniques < 0:
		msg = "record.uniques cannot be negative"
	}

	if msg != "" {
		return errors.New(msg)
	}
	return nil
}
