// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	product
	published interface{}
}

// book satisfies the fmt.Stringer
func (b *book) String() string {
	p := format(b.published)

	// product.String has a pointer receiver.
	// that's why we need to take its address here.
	//
	// when you put a value in an interface, the interface
	// cannot run that value's type's value-receiver methods.
	return fmt.Sprintf("%s - (%v)", &b.product, p)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"

	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
