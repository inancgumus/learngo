// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package magic

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Detect returns the files that have a valid header (file signature).
// A valid header is determined by the format.
func Detect(format string, filenames []string) (valids []string, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			err = fmt.Errorf("cannot detect: %v", rerr)
		}
	}()

	return detect(format, filenames), nil
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
