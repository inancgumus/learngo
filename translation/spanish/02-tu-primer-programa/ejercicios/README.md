1. **Imprime tu nombre y el nombre de tu mejor amigo** usando Println dos veces. [Échale un vistazo a este ejercicio](https://github.com/inancgumus/learngo/tree/master/translation/spanish/02-tu-primer-programa/ejericios/01-imprimir-nombres).

2. **Imprime tu GOPATH** usando la herramienta `go env` [Échale un vistazo a este ejercicio](https://github.com/inancgumus/learngo/tree/master/translation/spanish/02-tu-primer-programa/ejercicios/02-imprimir-gopath).

3. **Saludate a ti mismo.**

   1. Complila tu programa usando `go build`

   2. **Envíaselo a tu amigo**

      Debería tener el mismo sistema operativo.

      Por ejemplo si estás usando Windows envíaselo a un amigo que también tenga Windows.

   3. **Envíaselo a un amigo con un sistema operativo diferente**

      Deberías compilar el programa especificamente para su sistema operativo.

      **Crea un ejecutable para OS X:**
      `GOOS=darwin GOARCH=386 go build`

      **Crea un ejecutable para Windows:**
      `GOOS=windows GOARCH=386 go build`

      **Crea un ejecutable para Linux:**
      `GOOS=linux GOARCH=arm GOARM=7 go build`

      **Puedes encontrar una lista completa aquí:**
      https://golang.org/doc/install/source#environment

      **NOTA:** Si estás usando la consola de comandos o PowerShell, quizás necesites usarlo de la siguiente forma:
      `cmd /c "set GOOS=darwin GOARCH=386 && go build"`

4. **Llama a [Print](https://golang.org/pkg/fmt/#Print) en vez de a [Println](https://golang.org/pkg/fmt/#Println)** y observa que pasa.

5. **Llama a [Println](https://golang.org/pkg/fmt/#Println) o a [Print](https://golang.org/pkg/fmt/#Print) con multiples valores** separados por comas.

6. **Elimina las comillas dobles de una cadena de texto literal** y observa que pasa.

7. **Mueve tanto el package como import** al final del archivo y observa que pasa.

8. **[Lee la documentación en línea](https://golang.org/pkg)**.

   1. Échale un vistazo a qué son los packages y qué es lo que hacen.

   2. Mira su codigo fuente clickeando sobre su titulo.

   3. No tienes porque entenderlo todo, simplemente hazlo. Esto te dará un empujón de cara a las siguientes clases.

9. También, puedes tomar **un tour con Go**: https://tour.golang.org/

   1. Échale un vistazo rápido. Mira las caracteristicas del lenguaje.

   2. Hablaremos de ellas proximamente.
