// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// This program uses magic values

func main() {
	cm := 100
	m := cm / 100 // 100 is a magic value
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / 100 // 100 is a magic value
	fmt.Printf("%dcm is %dm\n", cm, m)
}
