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
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://inancgumus.github.com/x/rosie.unknown")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()

	// creates a file (or truncates if the file exists)
	file, err := os.Create("rosie.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	n, err := transfer(resp.Body, file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%d bytes transferred.\n", n)

	// main() runs the deferred funcs in reverse order:
	// file.Close()
	// resp.Body.Close()
}

// transfer streams data from a reader to a writer.
// if the reader doesn't contain a PNG signature, transfer returns errNotPNG.
// otherwise, it transfers the data to the writer.
func transfer(r io.Reader, w io.Writer) (n int64, err error) {
	const pngSign = "\x89PNG\r\n\x1a\n"
	const pngSignLen = 8

	var memory bytes.Buffer

	// limit what copy() reads.
	// lr := io.LimitReader(r, pngSignLen)
	// if n, err = io.Copy(&memory, lr); err != nil {
	// return n, err
	// }

	// same as above. behind the scenes CopyN() calls LimitReader.
	if n, err = io.CopyN(&memory, r, pngSignLen); err != nil {
		return n, err
	}

	// check the png signature.
	if !bytes.HasPrefix(memory.Bytes(), []byte(pngSign)) {
		return n, errors.New("not png")
	}

	// stitch the PNG signature (memory) and the response body reader together.
	// then copy them successively to the writer (*os.File).
	return io.Copy(w, io.MultiReader(&memory, r))
}
