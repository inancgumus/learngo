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
	"strconv"
	"time"
)

type book struct {
	product
	Published interface{}
}

func (b *book) String() string {
	p := format(b.Published)
	return fmt.Sprintf("%s - (%v)", &b.product, p)
}

func (b *book) MarshalJSON() ([]byte, error) {
	type jbook book

	jb := (*jbook)(b)
	jb.Published = format(b.Published)

	return json.Marshal(jb)
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

	const layout = "2006/01"
	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
