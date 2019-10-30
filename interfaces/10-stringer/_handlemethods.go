// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// In the depths of the Go standard library's fmt package...
// Printing functions use the handleMethods method.

// Example:
// var pocket money = 10
// fmt.Println(pocket)

//            the argument can be any type of value
//      stores the pocket variable in the argument variable
//                            ^
//                            |
func (p *pp) handleMethods(argument interface{}) (handled bool) {
	// ...

	// Checks whether a given argument is an error or an fmt.Stringer
	switch v := argument.(type) {
	// ...
	// If the argument is a Stringer, calls its String() method
	case Stringer:
		// ...
		//        pocket.String()
		//              ^
		//              |
		p.fmtString(v.String(), verb)
		return
	}

	// ...
}

/*
The original `handleMethods` code is more involved:

https://github.com/golang/go/blob/6f51082da77a1d4cafd5b7af0db69293943f4066/src/fmt/print.go#L574

	-> 574#handleMethods(..)
	-> 627#Stringer type check: If `v` is a Stringer, run:
	-> 630#v.String()
*/