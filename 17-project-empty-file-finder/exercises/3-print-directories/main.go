// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"log"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Find and write the names of subdirectories to a file
//
//  Create a program that can get multiple directory paths from
//  the command-line, and prints only their subdirectories into a
//  file named: dirs.txt
//
//
//  1. Get the directory paths from command-line
//
//  2. Append the names of subdirectories inside each directory
//     to a byte slice
//
//  3. Write that byte slice to dirs.txt file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Please provide directory paths
//
//   go run main.go dir/ dir2/
//
//   cat dirs.txt
//
//     dir/
//             subdir1/
//             subdir2/
//
//     dir2/
//             subdir1/
//             subdir2/
//             subdir3/
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + Get all the files in a directory using ioutil.ReadDir
//     (A directory is also a file)
//
//   + You can use IsDir method of a FileInfo value to detect
//     whether a file is a directory or not.
//
//     Check out its documentation:
//
//     go doc os.FileInfo.IsDir
//
//     # or using godocc
//     godocc os.FileInfo.IsDir
//
//   + You can use '\t' escape sequence for indenting the subdirs.
//
//   + You can find a sample directory structure under:
//     solution/ directory
//
// ---------------------------------------------------------

func main() {
	input_dir := os.Args[1:]

	if len(input_dir) == 0 {
		log.Fatal("Please provide directory paths")
	}
	root_dir := input_dir[0]
	all_dirs, error := os.ReadDir(root_dir)

	if error != nil {
		log.Fatal(error)
	}

	data := []byte{}
	for _, dir := range all_dirs {
		if dir.IsDir() == true {
			data = append(data, dir.Name()...)
			data = append(data, '/', '\n')
			sub_dirs, _ := os.ReadDir(root_dir + "/" + dir.Name())
			for _, subdir := range sub_dirs {
				if subdir.IsDir() == true {
					data = append(data, '\t')
					data = append(data, subdir.Name()...)
					data = append(data, '/', '\n')
				}
			}
			data = append(data, '\n')
		}
	}
	err := os.WriteFile("dirs.txt", data, 0644)

	if err != nil {
		log.Fatal(err)
	}

}
