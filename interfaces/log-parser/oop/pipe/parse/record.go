package parse

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/inancgumus/learngo/interfaces/log-parser/oop/pipe"
)

const fieldsLength = 4

// record stores fields of a log line.
type record struct {
	Domain  string
	Page    string
	Visits  int
	Uniques int
}

// Str gets a string field by name.
func (r record) Str(field string) string {
	switch field {
	case "domain":
		return r.Domain
	case "page":
		return r.Page
	}
	panic(fieldErr(field))
}

// Int gets an integer field by name.
func (r record) Int(field string) int {
	switch field {
	case "visits":
		return r.Visits
	case "uniques":
		return r.Uniques
	}
	panic(fieldErr(field))
}

// Sum the numeric fields with another record.
func (r record) Sum(other pipe.Record) pipe.Record {
	if other == nil {
		return r
	}
	r.Visits += other.(record).Visits
	r.Uniques += other.(record).Uniques
	return r
}

// UnmarshalText to a *record.
func (r *record) UnmarshalText(p []byte) (err error) {
	fields := strings.Fields(string(p))
	if len(fields) != fieldsLength {
		return fmt.Errorf("wrong number of fields %q", fields)
	}

	r.Domain, r.Page = fields[0], fields[1]

	if r.Visits, err = parseStr("visits", fields[2]); err != nil {
		return err
	}
	if r.Uniques, err = parseStr("uniques", fields[3]); err != nil {
		return err
	}
	return validate(*r)
}

// UnmarshalJSON to a *record.
func (r *record) UnmarshalJSON(data []byte) error {
	// `methodless` doesn't have any methods including UnmarshalJSON.
	// This trick prevents the stack-overflow (infinite loop).
	type methodless record

	var m methodless
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	// Cast back to the record and save.
	*r = record(m)

	return validate(*r)
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
	case r.Domain == "":
		err = errors.New("record.domain cannot be empty")
	case r.Page == "":
		err = errors.New("record.page cannot be empty")
	case r.Visits < 0:
		err = errors.New("record.visits cannot be negative")
	case r.Uniques < 0:
		err = errors.New("record.uniques cannot be negative")
	}
	return
}

func fieldErr(field string) error {
	return fmt.Errorf("record field: %q does not exist", field)
}
