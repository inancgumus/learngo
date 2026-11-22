// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// Licencia: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Para más tutoriales  : https://learngoprogramming.com
// Clases particulares  : https://www.linkedin.com/in/inancgumus/
// Sigueme en Twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello!")

	// Las declaraciones cambian el flujo de ejecución
	// Especialmente las declaraciones de control de flujo como "if"
	if 5 > 1 {
		fmt.Println("bigger")
	}
}