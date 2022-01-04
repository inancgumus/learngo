// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// Licencia: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Para más tutoriales  : https://learngoprogramming.com
// Clases particulares : https://www.linkedin.com/in/inancgumus/
// Sigueme en twitter: https://twitter.com/inancgumus

package main

import "fmt"

// No hable de esto en la lectura
// quiero dejarlo aqui como una pequeña nota
// Por favor revisalo.

var declarameOtraVez = 10

func anidado() { // Empieza el bloque de la función

	// declara las misma variables
	// ambas pueden existir juntas
	// Esta solo pertenece a esta función
	// la variable del paquete sigue intacta
	var declarameOtraVez = 5
	fmt.Println("Dentro del anidado:", declarameOtraVez)

} // Termina el bloque de la función

func main() { // Empieza el bloque de la función

	fmt.Println("Dentro de main:", declarameOtraVez)

	anidado()

	// a nivel paquete, declarameOtraVez no fue alterado
	// desde el cambio dentro de la funcion anidado
	fmt.Println("Dentro de main:", declarameOtraVez)

} // Termina el bloque de la función
