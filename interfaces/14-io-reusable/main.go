// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	n, err := transfer()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%d bytes transferred.\n", n)
}

func transfer() (n int64, err error) {
	// stream from the standard input to the in-memory buffer in 32KB data chunks.
	// os.Stdin.Read(...) -> os.Stdout.Write(...)
	if n, err = io.Copy(os.Stdout, os.Stdin); err != nil {
		return n, err
	}
	return n, err
}

// ioCopy streams from a file to another file.
// we use it to stream from the standard input to ouput.
func ioCopy(dst, src *os.File) error {
	// Use a fixed-length buffer to efficiently read from src stream in chunks.
	buf := make([]byte, 32768)

	// read over and over again to read all the data.
	for {
		// read can read only up to the buffer length.
		nr, er := src.Read(buf)
		// only write data if there is something to write.
		if nr > 0 {
			_, ew := dst.Write(buf[:nr])
			if ew != nil {
				return ew
			}
		}
		// io.EOF = there is nothing left to read—close the loop.
		if er == io.EOF {
			return nil
		}
		// Only return an error if the reading really fails.
		if er != nil {
			return er
		}
	}
	return nil
}
