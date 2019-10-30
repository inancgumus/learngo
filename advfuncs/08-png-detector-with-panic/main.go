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
	"os"
)

/*
	// ~~~ THE CLASSIC WAY ~~~
	try {
		// open a file
		// throws an exception
	} catch (ExceptionType name) {
		// handle the error
	} finally {
		// close the file
	}

	// ~~~ GO WAY ~~~
	file, err := // open the file
	if err != nil {
		// handle the error
	}
	// close the file

	// really really exceptional problem occurs
	// mostly due to a programmer

	// panic("PANIC PANIC")
*/

func main() {
	files := []string{
		"pngs/cups-jpg.png",
		"pngs/forest-jpg.png",
		"pngs/golden.png",
		"pngs/work.png",
		"pngs/shakespeare-text.png",
		"pngs/empty.png",
	}

	list("png", files)

	// fmt.Println("catch me if you can!")
}

func list(format string, files []string) {
	valids := detect(format, files)

	fmt.Printf("Correct Files:\n")
	for _, valid := range valids {
		fmt.Printf(" + %s\n", valid)
	}
}

func detect(format string, filenames []string) (valids []string) {
	header := headerOf(format)

	buf := make([]byte, len(header))

	for _, filename := range filenames {
		if read(filename, buf) != nil {
			continue
		}

		if bytes.Equal([]byte(header), buf) {
			valids = append(valids, filename)
		}
	}
	return
}

func headerOf(format string) string {
	switch format {
	case "png":
		return "\x89PNG\r\n\x1a\n"
	case "jpg":
		return "\xff\xd8\xff"
	}
	panic("unknown format: " + format)
}

// read reads len(buf) bytes to buf from a file
func read(filename string, buf []byte) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	if fi.Size() <= int64(len(buf)) {
		return fmt.Errorf("file size < len(buf)")
	}

	_, err = io.ReadFull(file, buf)
	return err
}
