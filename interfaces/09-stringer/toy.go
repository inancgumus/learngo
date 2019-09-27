// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

type toy struct {
	// toy is an fmt.Stringer
	// because the product is an fmt.Stringer
	// and the toy embeds the product
	product
}
