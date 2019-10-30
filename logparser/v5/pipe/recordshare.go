// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package pipe

import (
	"fmt"
	"reflect"
)

// Record stores the log line fields.
// The underlying fields are kept hidden.
// The users of the package are decoupled from the underlying record fields.
// When the fields change, the client won't feel the difference (at least in compile-time).
type Record struct {
	record
}

// Str returns a string field. Panics when the field doesn't exist.
func (r Record) Str(field string) string {
	return r.mustGet(field, reflect.String).String()
}

// Int returns an int field. Panics when the field doesn't exist.
func (r Record) Int(field string) int {
	return int(r.mustGet(field, reflect.Int).Int())
}

// Fields returns all the field names.
// The names can be used to query the Record.
func (r Record) Fields() (fields []string) {
	t := reflect.TypeOf(record{})

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		s := fmt.Sprintf("(%s: %s)", f.Name, f.Type.Name())

		fields = append(fields, s)
	}

	return
}

// mustGet the field with the same kind or panics.
func (r Record) mustGet(field string, kind reflect.Kind) reflect.Value {
	v := reflect.ValueOf(r.record).FieldByName(field)
	if !v.IsValid() {
		panic(fmt.Errorf("record.%s does not exist", field))
	}
	if v.Kind() != kind {
		panic(fmt.Errorf("record.%s is not %q", field, kind))
	}
	return v
}
