// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// Licencia: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Para más tutoriales  : https://learngoprogramming.com
// Clases particulares : https://www.linkedin.com/in/inancgumus/
// Sigueme en twitter: https://twitter.com/inancgumus

package main

func nope() { // Empieza el bloque de la función

	// hola y ok solamente son visibles aquí
	const ok = true
	var hola = "Hola"

	_ = hola
} // Termina el bloque de la función

func main() { // Empieza el bloque de la función

	// hola y ok no son visibles aquí

	// ERROR:
	// fmt.Println(hola, ok)

} // Termina el bloque de la función
