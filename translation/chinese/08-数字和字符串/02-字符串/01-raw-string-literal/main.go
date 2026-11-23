// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// The type of a string and a raw string literal
	// is the same. They both are strings.
	//
	// So, they both can be used as a string value.
	var s string
	s = "how are you?"
	s = `how are you?`
	fmt.Println(s)

	// string literal
	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	// raw string literal
	s = `
<html>
	<body>"Hello"</body>
</html>`

	fmt.Println(s)

	// windows path
	fmt.Println("c:\\my\\dir\\file") // string literal
	fmt.Println(`c:\my\dir\file`)    // raw string literal
}
