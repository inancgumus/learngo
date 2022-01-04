// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// Licencia: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Para más tutoriales  : https://learngoprogramming.com
// Clases particulares : https://www.linkedin.com/in/inancgumus/
// Sigueme en twitter: https://twitter.com/inancgumus

package main

import "fmt"

func hello() {
	// Solo este archivo puede acceder a la funcion de paquete fmt
	// cuando otros ya lo hiciero, ellos tamvien pueden acceder
	// Su propio 'fmt' "name"

	fmt.Println("hi! this is inanc!")
	bye()
}
