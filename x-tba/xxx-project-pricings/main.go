// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// Remember: This should point to the directory exactly after GOPATH
// Use / not \ even on Windows
import (
	"fmt"
)

func main() {
	// TODO: funcs can call other funcs

	const data = `New York,100,2,1,100000
Hong Kong,150,3,2,300000
Paris,200,4,3,250000
Istanbul,500,10,5,four hundred thousand`

	props, err := parse(data)
	if err != nil {
		fmt.Printf("> err: %s\n\n", err)
	}

	show(props)

	// fmt.Printf("%#v\n", props.list)
}
