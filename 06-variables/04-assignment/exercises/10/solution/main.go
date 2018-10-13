// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"path"
)

func main() {
	dir, _ := path.Split("secret/file.txt")

	fmt.Println(dir)
}
