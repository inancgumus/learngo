// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// You can easily write to a file using bash.
//
// However, when what you're creating is a library,
// then you won't have that option. So, it's good to learn
// how to write to a file using a byte slice.

// ioutil.ReadDir
// os.FileInfo

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

	// 1st: find the total size of all the empty files
	var total int
	for _, file := range files {
		if file.Size() == 0 {
			// +1 for the newline character
			// when printing the filename
			total += len(file.Name()) + 1
		}
	}

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
