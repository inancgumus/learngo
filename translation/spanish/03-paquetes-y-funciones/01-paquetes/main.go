// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// Licencia: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Para más tutoriales  : https://learngoprogramming.com
// Clases particulares : https://www.linkedin.com/in/inancgumus/
// Sigueme en twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	// Puedes acceder a funciones desde otros archivos
	// los cuales esten en el mismo paquete
	// Por ejemplo, `main()` puede acceder a `bye()` y `hey()`

	// Porque bye.go, hey.go y main.go
	//   estan en el mismo paquete.

	bye()
	hey()
}
