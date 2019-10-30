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
	paths := os.Args[1:]
	if len(paths) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	var dirs []byte

	for _, dir := range paths {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
			return
		}

		dirs = append(dirs, dir...)
		dirs = append(dirs, '\n')

		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, '\t')
				dirs = append(dirs, file.Name()...)
				dirs = append(dirs, '/', '\n')
			}
		}

		dirs = append(dirs, '\n')
	}

	err := ioutil.WriteFile("dirs.txt", dirs, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
