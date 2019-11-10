// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"errors"
	"io"
)

// create a reader for detecting the png signatures.
func pngReader(r io.Reader) io.Reader {
	return &signatureReader{r: r, sign: []byte("\x89PNG\r\n\x1a\n")}
}

// reader reads from `r` if the stream starts with PNG signature.
// otherwise it stops and returns with an error.
type signatureReader struct {
	r    io.Reader // reads from the response.Body (or from any reader)
	sign []byte    // stream should start with this initial signature
}

// Read implements the io.Reader interface.
func (sr *signatureReader) Read(b []byte) (n int, err error) {
	n, err = sr.r.Read(b)

	l := len(sr.sign)

	// simply return if the signature has already been detected.
	if l == 0 {
		return
	}
	// limit the compared bytes at most to the buffer length.
	if lb := len(b); l > lb {
		l = lb
	}
	if string(b[:l]) != string(sr.sign[:l]) {
		err = errors.New("not png")
	}
	// the next read will compare the rest of the signature after `l`
	// `sr.sign` will be zero after one or several reads later.
	sr.sign = sr.sign[l:]
	return
}
