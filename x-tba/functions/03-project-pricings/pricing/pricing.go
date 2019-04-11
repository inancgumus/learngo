// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pricing

const separator = ","

// New parses and returns a new Properties data
func New(data string) Properties {
	return Properties{parse(data)}
}
