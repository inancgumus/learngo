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
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}
	root_dir := args[0]
	files, err := os.ReadDir(root_dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		file_info, _ := f.Info()
		if file_info.Size() == 0 {
			fmt.Println("Empty_file: ", f.Name())
		}
	}
}
