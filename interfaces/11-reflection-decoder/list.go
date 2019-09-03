// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type item interface {
	discount(float64)
	fmt.Stringer
}

type schema struct {
	Category string
	Item     item
}

type list []item

func (l list) MarshalJSON() ([]byte, error) {
	var schemas []schema

	for _, it := range l {
		cat := reflect.TypeOf(it).Elem().Name()
		schemas = append(schemas, schema{cat, it})
	}

	return json.Marshal(schemas)
}

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚."
	}

	var str strings.Builder

	for _, it := range l {
		str.WriteString(it.String())
		str.WriteRune('\n')
	}

	return str.String()
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		it.discount(ratio)
	}
}
