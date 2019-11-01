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
	"log"
	"os"
)

func main() {
	// Stdin, Stdout, and Stderr are *os.File.

	// #1: reads from the standard input
	// ------------------------------------------------
	// if err := read(os.Stdin); err != nil {
	// 	log.Fatalln(err)
	// }

	// #2: writes to the standard output
	// ------------------------------------------------
	// buf := []byte("hi!\n")
	// if err := write(os.Stdout, buf); err != nil {
	// 	log.Fatalln(err)
	// }

	// #2b: reads the entire file content into memory.
	// ------------------------------------------------
	// buf, err := ioutil.ReadFile("alice.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// if err := write(os.Stdout, buf); err != nil {
	// 	log.Fatalln(err)
	// }

	// #3: reads and writes with 32KB of memory
	//     no matter how large the source data is.
	// ------------------------------------------------
	if err := ioCopy(os.Stdout, os.Stdin); err != nil {
		log.Fatalln(err)
	}
}

func ioCopy(dst, src *os.File) error {
	buf := make([]byte, 32768)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			_, ew := dst.Write(buf[:nr])
			if ew != nil {
				return ew
			}
		}
		if er == io.EOF {
			return nil
		}
		if er != nil {
			return er
		}
	}
	return nil
}

func write(dst *os.File, buf []byte) error {
	nw, ew := dst.Write(buf)

	fmt.Printf("wrote    : %d bytes\n", nw)
	fmt.Printf("write err: %v\n", ew)

	return ew
}

func read(src *os.File) error {
	// Use a fixed-length buffer to efficiently read from src stream in chunks.
	buf := make([]byte, 32768)

	// #1b: read over and over again to read all the data.
	for {
		// #1a: read can read only up to the buffer length.
		nr, er := src.Read(buf)

		fmt.Printf("\n")
		// fmt.Printf("buf      : %q\n", buf[0:nr])
		fmt.Printf("read     : %d bytes\n", nr)
		fmt.Printf("read err : %v\n", er)

		// #1c: io.EOF = there is nothing left to read—close the loop.
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
