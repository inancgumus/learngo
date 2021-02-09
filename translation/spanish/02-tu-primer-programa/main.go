// Copyright © 2021 Inanc Gumus
// Learn Go Programming Course
// Licencia: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Para mas tutoriales: https://learngoprogramming.com
// Clases particulares: https://www.linkedin.com/in/inancgumus/
// Sigueme en twitter: https://twitter.com/inancgumus

// package main es un paquete especial
// le permite a Go crear un archivo ejecutable
package main

/*
Esto es un comentario multilinea.

la palabra import hace que otro paquete este disponible
  para este "archivo" .go.

import "fmt" te permite acceder a la funcionaliddad del paquete fmt
  en este archivo.
*/
import "fmt"

// "func main" es especial.
//
// Go tiene que saber por donde empezar
//
// func main crea un punto de inicio para Go
//
// Despues de compilar el codigo,
// Go ejecutara esta funcion primero.
func main() {
	// despues del: import "fmt"
	// La funcion Println del paquete "fmt" estara disponible

	// Lee sobre ella escribiendo los siguiente en la consola:
	//   go doc -src fmt Println

	// Println es simplemente una funcion exportada de:
	//   "fmt" package

	// Exportada = Primera letra en mayuscula
	fmt.Println("Hello Gopher!")

	// Go no puede llamar a la funcion Println por si mismo.
	// Por eso la tienes que llamar aqui.
	// Solo llama a la `func main` de forma automatica.

	// -----

	// Go soporta caracteres unicode en cadenas de texto literal
	// y tambien en el codigo fuente: KÖSTEBEK!
	//
	// Porque: Literal ~= Codigo Fuente
}
