// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// Remember: This should point to the directory exactly after GOPATH
// Use / not \ even on Windows
import "github.com/inancgumus/learngo/x-tba/3-functions/xxx-project-pricings/pricing"

func main() {
	// TODO: funcs can call other funcs

	const data = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

	props := pricing.New(data)
	pricing.Print(props)

	// fmt.Printf("%#v\n", props.list)
}
