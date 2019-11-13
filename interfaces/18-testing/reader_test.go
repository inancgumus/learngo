package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestCorrectSignature(t *testing.T) {
	const signature = "\x89PNG\r\n\x1a\n"
	const input = signature + "REST OF THE DATA"

	pr := pngReader(strings.NewReader(input))

	out, err := ioutil.ReadAll(pr)
	if err != nil {
		t.Fatalf("got: %q; want: nil err", err)
	}
	if string(out) != input {
		t.Fatalf("got: % x; want: % x", out, input)
	}
}

func TestIncorrectSignature(t *testing.T) {
	const input = "WITHOUT PNG SIGNATURE"

	// strings.Reader is a read-only, in-memory stream reader.
	// It's an io.Reader implementation.
	pr := pngReader(strings.NewReader(input))

	// `ioutil.ReadAll` is a dangerous but useful utility function
	// that can read all the data from a reader into memory.
	// Don't use it for large (?) data.
	_, err := ioutil.ReadAll(pr)
	if err == nil {
		t.Fatalf("got: %v; want: !nil err", err)
	}
}

func TestSummation(t *testing.T) {
	t.Skip() // skip this test

	result := 1 + 2
	if result != 5 { // it'll always fail
		// t.Fatal("want: 5") // not a good error message
		t.Fatalf("got: %d; want: 5", result) // explanatory error message: good.
	}
}
