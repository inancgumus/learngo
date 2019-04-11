package pricing

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	const data = `New York,100,2,1,100000
Istanbul,150,3,2,200000`

	props := parse(data)

	if l := len(props); l != 2 {
		t.Fatalf(`got: %d; want: 2`, l)
	}

	want := Property{
		Location: "New York",
		Size:     100, Beds: 2, Baths: 1, Price: 100000,
	}
	p := props[0]

	if !reflect.DeepEqual(want, p) {
		t.Errorf(`got: %#v; want: %#v`, p, want)
	}

	want = Property{
		Location: "Istanbul",
		Size:     150, Beds: 3, Baths: 2, Price: 200000,
	}

	p = props[1]
	if !reflect.DeepEqual(want, p) {
		t.Errorf(`got: %#v; want: %#v`, p, want)
	}
}

func TestParseBetter(t *testing.T) {
	const data = `New York,100,2,1,100000
Istanbul,150,3,2,200000`

	tests := []Property{
		{
			Location: "New York",
			Size:     100, Beds: 2, Baths: 1, Price: 100000,
		},
		{
			Location: "Istanbul",
			Size:     150, Beds: 3, Baths: 2, Price: 200000,
		},
	}

	props := parse(data)

	if l := len(props); l != 2 {
		t.Fatalf(`got: %d; want: 2`, l)
	}

	for i, p := range tests {
		if !reflect.DeepEqual(p, props[i]) {
			t.Errorf(`got: %#v; want: %#v`, props[i], p)
		}
	}
}
