# Escribe Tu Primer Programa Usando Go

Hola!

Puedes guardar este documento despues de ver los videos de esta sección en particular.

Tambien puedes imprimirla y usarlo junto a los videos de esta sección.

Disfruta!

---

## COMANDOS DE LA CONSOLA DE COMANDOS 

* **WINDOWS:**

* Entra al siguiente directorio: `cd directoryPath`

    * Muestra los archivos en ese directorio: `dir`

* **OS X & LINUXes:**

    * Muestra los archivos en ese directorio: `ls`

## COMPILANDO Y EJECUTANDO PROGRAMAS CON GO:

* **Compilando Un Programa En Go:**

    * Dentro del directorio del programa, escribe:
        * `go build`

* **Ejecuta Un Programa En Go:**

    * Dentro del directorio del programa, escribe:
        * `go run main.go`

## QUE ES EL $GOPATH?

* _$GOPATH_  es una variable de entorno que apunta a un directorio donde los archivos descargados por Go y los tuyos propios se encuentran.

    * **En Windows**, Esta en: `%USERPROFILE%\go`

    * **En OS X y en Linux**, Esta en: `~/go`

    * **NOTE:** Nunca introduzcas tu `GOPATH` de forma manual. Siempre se encuentra en el directorio de usuario por defecto.

* **GOPATH tiene tres directorios:**

    * **src:** Contiene los archivos fuente para tus paquetes o para otros paquetes que descargues. Puedes compilar y ejecutar programas mientras te encuentres en el directorio del programa.

    * **pkg:** Contiene paquetes de archivos compilados. Go usa esto para compilar de forma eficiente.

    * **bin:** Contiene programas compilados y ejecutables. Cuando usas el comando go install en el directorio de un programa, Go creara un archivo ejecutable en el mismo.

        * _Quizas quieraas añadir esto a tu variable de entorno`PATH` si no esta ahi._

## DONDE DEBERÍA PONER MIS ARCHIVOS FUENTE?

* `$GOPATH/src/github.com/yourUsername/programDirectory`

* **Ejemplo:**

    * Mi nombre en github es: inancgumus

    * Asi que añado todos mis programas bajo: `~/go/src/github.com/inancgumus/`

    * Asi que, digamos que tengo un programa llamado hello, y lo pongo bajo este directorio: `~/go/src/github.com/inancgumus/hello`

## TU PRIMER PROGRAMA!

* **Crea Los Directorios:**
    * **OS X & Linux (o git bash):**
        * Crea un nuevo directorio:
            * `mkdir -p ~/go/src/yourname/hello`
        * Ves a ese directorio:
            * `cd ~/go/src/yourname/hello`

    * **Windows:**
        * Crea un nuevo directorio:
            * `mkdir c:\Go\src\yourname\hello`
        * Ves a ese directorio:
            * `cd c:\Go\src\yourname\hello`

* Crea un nuevo archivo usando `code main.go` Visual Studio Code.
* Añade el siguiente codigo y guardalo.
* Vuelve a la consola de comandos.
    * Ejecutalo con este comando: `go run main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi! I want to be a Gopher!")
}
```

Eso es todo! Disfruta!

## NOTA:

* Hay un nuevo *Go Modules* que te permite ejecutar tus programas en cualquier directorio de tu elección. Es todavia una caracteristica experimental, cuando se estabilize actualizare el curso para incluir los Go Modules.

> Para mas tutoriales: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2021 Inanc Gumus
>
> Learn Go Programming Course
>
> [Click aqui para leer la licencia.](https://creativecommons.org/licenses/by-nc-sa/4.0/)

