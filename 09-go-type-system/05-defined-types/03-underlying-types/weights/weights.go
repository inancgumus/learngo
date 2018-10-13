// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package weights

type (
	// Gram underlying type is int64
	Gram int

	// Kilogram underlying type is int64
	Kilogram Gram

	// Ton underlying type is int64
	Ton Kilogram
)
