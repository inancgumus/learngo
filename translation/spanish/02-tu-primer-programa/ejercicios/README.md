1. **Imprime tu nombre y el nombre de tu mejor amigo** usando Println dos veces. [Echale un vistazo a este ejercicio](https://github.com/inancgumus/learngo/tree/master/translation/spanish/02-tu-primer-programa/ejericios/01-imprimir-nombres).

2. **Imprime tu GOPATH** usando la herramienta `go env` [Echale un vistazo a este ejercicio](https://github.com/inancgumus/learngo/tree/master/translation/spanish/02-tu-primer-programa/ejercicios/02-imprimir-gopath).

3. **Saludate a ti mismo.**

    1. Complila tu programa usando `go build`

    2. **Enviaselo a tu amigo**

       Deberia tener el mismo sistema operativo.

       Por ejemplo si estas usando Windows enviaselo a un amigo que tambien tenga Windows.

    3. **Enviaselo a un amigo con un sistema operativo diferente**

       Deberias compilar el programa especificamente para su sistema operativo.

       **Crea un ejecutable para OS X:**
       `GOOS=darwin GOARCH=386 go build`

       **Crea un ejecutable para Windows:**
       `GOOS=windows GOARCH=386 go build`

       **Crea un ejecutable para Linux:**
       `GOOS=linux GOARCH=arm GOARM=7 go build`

       **Puedes encontrar una lista completa aqui:**
       https://golang.org/doc/install/source#environment

       **NOTA:** Si estas usando la consola de comandos o PowerShell, quizas necesites usarlo de la siguiente forma:
       `cmd /c "set GOOS=darwin GOARCH=386 && go build"`

4. **Llama a [Print](https://golang.org/pkg/fmt/#Print) en vez de a [Println](https://golang.org/pkg/fmt/#Println)** y observa que pasa.

5. **Llama a [Println](https://golang.org/pkg/fmt/#Println) o a [Print](https://golang.org/pkg/fmt/#Print) con multiples valores** separados por comas.

6. **Elimina las comillas dobles de una cadena de texto literal** y observa que pasa.

7. **Mueve tanto package como import** al final del archivo y observa que pasa.

8. **[Lee la documentación en linea](https://golang.org/pkg)**.

    1. Echale un vistazo a que son los packages y que es lo que hacen.

    2. Mira su codigo fuente dandole click a su titulo.

    3. No tienes porque entenderlo todo, simplemente hazlo. Esto te dara un empujon de cara a las siguientes clases.

9. Tambien, puedes tomar **un tour con Go**: https://tour.golang.org/

    1. Echale un vistazo rapido. Mira las caracteristicas del lenguaje.

    2. Hablaremos de ellas proximamente.
