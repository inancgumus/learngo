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
	//Dos archivos pertenecen al mismo paquete
	// llamando `bye()` de bye.go a aqui
	bye()
}

// Ejercicio: Quita los comentarios de la siguiente función y analiza el error

// func bye() {
// 	fmt.Println("Bye!")
// }

// ***** EXPLICACIÓN *****
//
// ERROR: "bye" function "redeclared"
//        in "this block"
//
// "this block" significa = "main package"
//
// "redeclared" significa = Estas usando el mismo nompre en la funcion otra vez
//
// La función main package es:
// todo el codigo que esta en el mismo paquete main
