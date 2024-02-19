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
	"log"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func main() {
	items := os.Args[1:]

	if len(items) == 0 {
		log.Fatal("Send me some items and I will sort them")
	}

	sort.Strings(items)

	var total_size int

	for _, item := range items {
		total_size += len(item) + 1
	}

	file_data := make([]byte, 0, total_size)

	for _, item := range items {
		file_data = append(file_data, item...)
		file_data = append(file_data, '\n')
	}

	err := os.WriteFile("sorted.txt", file_data, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", file_data)

}
