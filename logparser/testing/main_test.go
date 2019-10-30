// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// +build integration

// go test -tags=integration

package main_test

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

const (
	okIn = `
a.com 1 2
b.com 3 4
a.com 4 5
b.com 6 7`

	okOut = `
DOMAIN                             VISITS           TIME SPENT
-----------------------------------------------------------------
a.com                                   5                    7
b.com                                   9                   11

TOTAL                                  14                   18`
)

func TestSummary(t *testing.T) {
	tests := []struct {
		name, in, out string
	}{
		{"valid input", okIn, okOut},
		{"missing fields", "a.com 1 2\nb.com 3", "> Err: line #2: missing fields: [b.com 3]"},
		{"incorrect visits", "a.com 1 2\nb.com -1 1", `> Err: line #2: incorrect visits: "-1"`},
		{"incorrect time spent", "a.com 1 2\nb.com 3 -1", `> Err: line #2: incorrect time spent: "-1"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(t, strings.TrimSpace(tt.in), strings.TrimSpace(tt.out))
		})
	}
}

func run(t *testing.T, in, out string) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = strings.NewReader(in)

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(got, []byte(out+"\n")) {
		t.Fatalf("\nwant:\n%s\n\ngot:\n%s", out, got)
	}
}
