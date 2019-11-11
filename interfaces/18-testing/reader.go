// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bytes"
	"fmt"
	"io"
)

// reader reads from `r` if the stream starts with a given signature.
// otherwise it stops and returns with an error.
type signatureReader struct {
	r         io.Reader // reads from the response.Body (or from any reader)
	signature []byte    // stream should start with this initial signature
}

// Read implements the io.Reader interface.
func (sr *signatureReader) Read(b []byte) (n int, err error) {
	n, err = sr.r.Read(b)

	l := len(sr.signature)
	if l == 0 {
		return
	}
	if lb := len(b); lb < l {
		l = lb
	}
	if got, want := b[:l], sr.signature[:l]; !bytes.Equal(got, want) {
		err = fmt.Errorf("signature doesn't match, got: % x, want: % x", got, want)
	}
	sr.signature = sr.signature[l:]
	return n, err
}

// create a reader for detecting only the png signatures.
func pngReader(r io.Reader) io.Reader {
	return &signatureReader{
		r:         r,
		signature: []byte("\x89PNG\r\n\x1a\n"),
	}
}
