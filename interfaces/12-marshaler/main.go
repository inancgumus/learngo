// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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
  - It cannot encode custom types.
  - Implement the json.Marshaler interface and teach it how to encode your custom types.
  - See time.Time code: It satisfies the json.Marshaler interface.

- json.Unmarshal() can only decode primitive types.
  - It cannot decode custom types.
  - Implement the json.Unmarshaler interface and teach it how to decode your custom types.
  - See time.Time code: It satisfies the json.Unmarshaler interface as well.

- strconv.AppendInt() can append an int value to a []byte.
  - There are several other functions in the strconv package for other primitive types as well.
  - Do not make unnecessary string <-> []byte conversions.

- log.Fatalln() can print the given error message and terminate the program.
  - Use it only from the main(). Do not use it in other functions.
  - main() is should be the main driver of a program.
*/
