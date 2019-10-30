// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"log"
)

func main() {
	// First encode products as JSON:
	data, err := encode()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))

	// Then decode them back from JSON:
	l, err := decode(data)
	if err != nil {
		log.Fatalln(err)
	}

	// Let the list value print itself:
	fmt.Print(l)
}

/*
Summary:

- json.Marshal() and json.MarshalIndent() can only encode primitive types.
  - Custom types can tell the encoder how to encode.
  - To do that satisfy the json.Marshaler interface.

- json.Unmarshal() can only decode primitive types.
  - Custom types can tell the decoder how to decode.
  - To do that satisfy the json.Unmarshaler interface.

- strconv.AppendInt() can append an int value to a []byte.
  - There are several other functions in the strconv package for other primitive types as well.
  - Do not make unnecessary string <-> []byte conversions.

- log.Fatalln() can print the given error message and terminate the program.
  - Use it only from the main(). Do not use it in other functions.
  - main() is should be the main driver of a program.
*/
