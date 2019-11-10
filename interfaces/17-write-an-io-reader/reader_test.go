package main

import (
	"io"
	"strings"
	"testing"
)

func Test_PNG_Reader_Correct_Signature(t *testing.T) {
	const input = "\x89PNG\r\n\x1a\nHELLO"

	out, err := readSignature(input)
	if err != nil {
		t.Fatalf("got: %q; want: <nil> err", err)
	}
	if out != input {
		t.Fatalf("invalid output, got: '%x'; want: '%x'", out, input)
	}
}

func Test_PNG_Reader_Incorrect_Signature(t *testing.T) {
	_, err := readSignature("\x89INCORRECT")
	if err == nil {
		t.Fatal("got: nil; want: !nil err")
	}
}

func readSignature(in string) (string, error) {
	r := strings.NewReader(in)
	w := strings.Builder{}
	_, err := io.Copy(&w, pngReader(r))
	return w.String(), err
}
