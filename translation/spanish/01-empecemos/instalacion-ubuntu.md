# Instalación Linux (Ubuntu)

Si quieres usar snap puedes instalar Go facilmente con este comando:

```bash
sudo snap install go --classic
```

Sino, puedes seguir los siguientes pasos

## 1. Actualiza los paquetes locales

```bash
sudo apt-get update
```

## 2. Instala Git

```bash
sudo apt-get install git
```

## 3. Instala Go

Hay dos formas:

1. Desde la pagina web: selecciona linux y empezara la descarga.

    ```bash
    firefox https://golang.org/dl
    ```

2. Si estas usando snap: avanza hasta el paso 5.

    ```bash
    sudo snap install go --classic
    ```

## 4. Copia Go en el directorio apropiado

1. Encuentra el nombre del archivo descargado

2. Usa el nombre de ese archivo para descomprimirlo

    ```bash
    gofile="BORRA_ESTO_Y_ESCRIBE_EL_NOMBRE_DEL_ARCHIVO_DESCARGADO (sin la extension)"
    tar -C /usr/local -xzf ~/Downloads/$gofile 
    ```

## 5. Añade los executables de Go a tu PATH

1. Añade el directorio `go/bin` al `$PATH` para poder ejecutar los comandos principales de Go.

    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    ```

2. Añade el directorio `$HOME/go/bin` a tu `$PATH`

    ```bash
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.profile
    ```

## Instala Las Herramientas de Go

* Estas son utiles herramientas que ayudan a facilitar el desarrollo (como goimports)

* `go get` no puede ser usado sin haber instalado un programa de control de versiones como Git.

* Esto creara un directorio llamado `~/go` donde se descargaran las herramientas

* Este directorio tambien es el lugar donde deberias poner tu codigo. (Si no vas a usar los modulos de Go)

    ```bash
    go get -v -u golang.org/x/tools/...
    ```

## Instala El Editor Visual Studio Code

Nota: Puedes usar otro editor de codigo si quieres. Igualmente, el curso usa Visual Studio Code (VS Code).

1. Abre la aplicación "Ubuntu Software"
2. Busca "VSCode" y clicka en el boton de Instalar

## PASO OPCIONAL

1. Crea un archivo llamado `hola.go` en un nuevo directorio alejado de `$GOPATH`

    ```bash
    cat <<EOF > hola.go
    package main

    import "fmt"

    func main() {
        fmt.Println("hola gopher!")
    }
    EOF
    ```

2. Ejecuta el programa

    ```bash
    go run hola.go
    Deberia imprimir: hola gopher!
    ```

## Eso es todo! Disfrutad! 🤩

<div style="page-break-after: always;"></div>

> Para mas tutoriales: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2021 Inanc Gumus
> 
> Learn Go Programming Course
> 
> [Click aqui para leer la licencia.](https://creativecommons.org/licenses/by-nc-sa/4.0/)