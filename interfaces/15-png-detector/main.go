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
	"net/http"
	"os"
)

func main() {
	// initiate the transmission channel (http connection) to the webserver.
	resp, err := http.Get("https://inancgumus.github.com/x/rosie.unknown")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	// close it so the http package can reuse the connection.
	defer resp.Body.Close()

	// resp.Body here is an io.ReadCloser: Read() + Close() methods.
	// but in the transfer function it's an io.Reader: Only the Read() method.
	n, err := transfer(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%d bytes transferred.\n", n)

	// resp.Body.Close() runs here (because of the defer above)
}

func transfer(r io.Reader) (n int64, err error) {
	const pngSign = "\x89PNG\r\n\x1a\n"
	const pngSignLen = 8

	// create an in-memory buffer (bytes.Buffer is an io.Reader and an io.Writer).
	var memory bytes.Buffer

	// stream from the standard input to the in-memory buffer in 32KB data chunks.
	// r.Read(...) -> memory.Write(...)
	if n, err = io.Copy(&memory, r); err != nil {
		return n, err
	}

	// check the png signature.
	buf := memory.Bytes()
	fmt.Printf("% x\n", buf[:pngSignLen])
	fmt.Printf("% x\n", pngSign)

	return n, err
}
