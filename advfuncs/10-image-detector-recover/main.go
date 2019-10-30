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

func main() {
	defer func() {
		if rerr := recover(); rerr != nil {
			fmt.Println("panicked:", rerr)
		}
	}()

	files := []string{
		"pngs/cups-jpg.png",
		"pngs/forest-jpg.png",
		"pngs/golden.png",
		"pngs/work.png",
		"pngs/shakespeare-text.png",
		"pngs/empty.png",
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("png or jpg?")
		return
	}

	list(args[0], files)

	// fmt.Println("catch me if you can!")
}

func list(format string, files []string) {
	valids := detect(format, files)

	fmt.Printf("Correct Files:\n")
	for _, valid := range valids {
		fmt.Printf(" + %s\n", valid)
	}
}

func detect(format string, filenames []string) (pngs []string) {
	header := headerOf(format)
	buf := make([]byte, len(header))

	for _, filename := range filenames {
		if read(filename, buf) != nil {
			continue
		}

		if bytes.Equal([]byte(header), buf) {
			pngs = append(pngs, filename)
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
	// this should never occur
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
