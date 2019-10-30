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
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("run with a PNG file")
		return
	}

	// this is not the best way
	// it's better only to read the first 24 bytes
	// this reads the whole file into memory
	img, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	report()
	img = append([]byte(nil), img[:24]...)
	// img = img[:24:24] // unnecessary
	report()

	// s.PrintBacking = true
	// s.MaxPerLine = 8
	// s.MaxElements = 24
	// s.PrintBytesHex = true
	// s.Show("first 24 bytes", img)

	var (
		pngHeader     = [...]byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}
		ihdrChunkType = [...]byte{0x49, 0x48, 0x44, 0x52}
	)

	// ------------------------------------------
	// Read the PNG header
	// https://www.w3.org/TR/2003/REC-PNG-20031110/#5PNG-file-signature
	// ------------------------------------------
	lpng := len(pngHeader)
	if !bytes.Equal(img[:lpng], pngHeader[:]) {
		fmt.Println("missing PNG header")
		return
	}

	// skip the png header
	img = img[lpng:]

	// ------------------------------------------
	// Read the IHDR chunk header
	// The IHDR chunk shall be the first chunk in the PNG datastream.
	// https://www.w3.org/TR/2003/REC-PNG-20031110/#11IHDR
	// ------------------------------------------
	header := img[:8] // get the length and chunk type

	// ensure that the chunk type is IHDR
	if !bytes.Equal(header[4:8], ihdrChunkType[:]) {
		fmt.Println("missing IHDR chunk")
		return
	}

	// ------------------------------------------
	// Read the IHDR Chunk Data
	// ------------------------------------------
	img = img[len(header):] // skip the IHDR chunk header data
	ihdr := img[:8]         // read the width&height from the ihdr chunk

	// All integers that require more than one byte shall be in:
	// network byte order = Big Endian
	// https://www.w3.org/TR/2003/REC-PNG-20031110/#7Integers-and-byte-order
	fmt.Printf("dimensions: %dx%d\n",
		// read the first 4 bytes (width)
		binary.BigEndian.Uint32(ihdr[:4]),
		// read the next 4 bytes (height)
		binary.BigEndian.Uint32(ihdr[4:8]))
}

func report() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf(" > Memory Usage: %v KB\n", m.Alloc/1024)
}
