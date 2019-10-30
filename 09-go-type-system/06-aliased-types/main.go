// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func main() {
	// aliased types are the same types
	// just with different names
	// for readability and refactoring
	var (
		// type byte = int8
		byteVal  byte
		uint8Val uint8
		intVal   int
	)

	uint8Val = byteVal // ok

	var (
		// type rune = int32
		runeVal  rune
		int32Val int32
	)

	runeVal = int32Val // ok

	runeVal = rune(int32Val)
	runeVal = rune(runeVal)

	// keep the compiler happy
	_, _, _, _ = byteVal, uint8Val, intVal, runeVal
}

// For the curious - compiler internals:
// https://github.com/golang/go/blob/4f1f503373cda7160392be94e3849b0c9b9ebbda/src/cmd/compile/internal/gc/universe.go#L409
