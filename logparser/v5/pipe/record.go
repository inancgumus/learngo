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

// Sum the numeric fields with another record.
func (r Record) Sum(other Record) Record {
	r.visits += other.visits
	r.uniques += other.uniques
	return r
}

// UnmarshalText to a *Record.
func (r *Record) UnmarshalText(p []byte) (err error) {
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

// UnmarshalJSON to a *Record.
func (r *Record) UnmarshalJSON(data []byte) error {
	// `methodless` doesn't have any methods including UnmarshalJSON.
	// This trick prevents the stack-overflow (infinite loop).
	type methodless Record

	var m methodless
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	// Cast back to the Record and save.
	*r = Record(m)

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
func validate(r Record) (err error) {
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
