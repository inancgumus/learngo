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
	"log"
	"os"
)

func main() {
	// you can download the rosie.unknown image in the link:
	// https://inancgumus.github.com/x/rosie.unknown

	// then feed the file to the standard input of this program:
	// go run . < rosie.unknown

	// create an in-memory buffer (bytes.Buffer is an io.Reader and an io.Writer).
	memory := bytes.Buffer{}

	// stream from the standard input to the in-memory buffer in 32KB data chunks.
	// os.Stdin.Read(...) -> memory.Write(...)
	if _, err := io.Copy(&memory, os.Stdin); err != nil {
		log.Fatal(err)
	}

	// get the accumulated bytes from the in-memory buffer.
	buf := memory.Bytes()

	// print the first eight bytes.
	fmt.Printf("% x\n", buf[:8])

	// compare it with the png signature.
	const pngSign = "\x89PNG\r\n\x1a\n"
	fmt.Printf("% x\n", pngSign)
}
