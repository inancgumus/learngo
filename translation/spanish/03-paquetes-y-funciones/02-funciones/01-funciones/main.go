// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// Licencia: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Para más tutoriales  : https://learngoprogramming.com
// Clases particulares : https://www.linkedin.com/in/inancgumus/
// Sigueme en twitter: https://twitter.com/inancgumus

package main

// Librerias importadas
import "fmt"

// variables del paquete
const ok = true

// función del paquete
func main() { // Empieza el bloque de la función
	var hola = "Hola"

	// hola y ok son visibles aquí
	fmt.Println(hola, ok)

} // Termina el bloque de la función
