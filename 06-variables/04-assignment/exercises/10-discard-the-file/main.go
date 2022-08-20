package main

import (
	"fmt"
	"path"
)

// ---------------------------------------------------------
// EXERCISE: Discard The File
//
//  1. Print only the directory using `path.Split`
//
//  2. Discard the file part
//
// RESTRICTION
//  Use short declaration
//
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

func main() {
	// UNCOMMENT THE CODE BELOW:

	dir, _ := path.Split("secret/file.txt")

	fmt.Println(dir)
}
