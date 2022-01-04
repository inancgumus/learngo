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
	// Como puedes verm no necesito importar el un paquete
	// y puedo llamar a la funcion 'hello' aqui
	//
	// Esto es porque las funciones de paquetes (librerias)
	// estan compartidas en el mismo paquete
	hello()
	// pero aqui no puedo acceder al paquete fmt sin importarlo
	//
	// Esto es porque esta en la funcion del archivo printer.go
	// lo importa
	//

	// esta funcion main puede tambien llamar a la funcion bye aqui
	// bye()
}

// printer.go puede llamar a esta funcion
//
// Esto es porque la funcion bye esta en la funcion del paquete (libreria)
// del paquete main ahora
// .
//
// La funcion main tambien puede llamar aqui.
func bye() {
	fmt.Println("bye bye")
}
