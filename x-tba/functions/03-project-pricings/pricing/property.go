// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pricing

// Property stores the data about a property
type Property struct {
	Location                 string
	Size, Beds, Baths, Price int
}

// Properties stores all the properties
type Properties struct {
	list []Property
}
