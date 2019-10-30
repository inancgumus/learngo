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
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 1st: You can also take the average of the total file length
	//      across platforms. It's about 255.
	//
	// https://en.wikipedia.org/wiki/Comparison_of_file_systems#Limits
	//
	// BTRFS   255 bytes
	// exFAT   255 UTF-16 characters
	// ext2    255 bytes
	// ext3    255 bytes
	// ext3cow 255 bytes
	// ext4    255 bytes
	// FAT32   255 bytes
	// NTFS    255 characters
	// XFS     255 bytes
	//
	// total := len(files) * 256

	// 1st B: To be exact, find the total size of all the empty files
	var total int
	for _, file := range files {
		if file.Size() == 0 {
			// +1 for the newline character
			// when printing the filename afterward
			total += len(file.Name()) + 1
		}
	}
	fmt.Printf("Total required space: %d bytes.\n", total)

	// 2nd: allocate a large enough byte slice in one go
	names := make([]byte, 0, total)

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()

			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", names)
}

// See: https://www.tutorialspoint.com/unix/unix-file-permission.htm
// See: http://permissions-calculator.org/
