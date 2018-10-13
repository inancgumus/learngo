// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// this one uses a raw string literal

	// can you see how readable it is?
	// compared to the previous one?

	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json)
}
