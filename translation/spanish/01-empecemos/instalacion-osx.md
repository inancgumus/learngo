# Instalación OSX

## NOTA

Si tienes [homebrew](https://brew.sh) instalado, puedes instalar Go facilmente de la siguiente forma:

```bash
# si no tienes git instalado instalalo asi
brew install git

# despues instala go
brew install go

# añade la ruta GOBIN a tu ruta en ~/.bash_profile
export PATH=${HOME}/go/bin:$PATH
```

## 1. Instala El Editor Visual Studio Code

1. Instalalo pero no lo abras todavia
2. Ves a [https://code.visualstudio.com](https://code.visualstudio.com)
3. Selecciona la version OS X (Mac) y empieza la descarga
4. Descomprime los archivos descargados y muevelos a la carpeta `~/Applications`

## 2. Instala Git

1. Descarga y ejecuta el instalador. Ves a: [https://git-scm.com/downloads](https://git-scm.com/downloads)
2. Selecciona VSCode como editor por defecto

## 3. Instala Go

1. Ves a [https://golang.org/dl](https://golang.go/dl)
2. Elige la versión OS X (Mac)
3. Ejecuta el instalador

## 4. Configura Visual Studio Code

1. Abre VS Code; haz click en la pestaña de extensiones situada a la izquierda, busca "go" e instalalo
2. Cierra VS Code por completo y ábrelo de nuevo

3. Ve al menu Ver: selecciona **Paleta de Comandos**
    1. O simplemente presiona: `ctrl+shift+p`
    2. Escribe: `go install`
    3. Elige: _"Go: Install/Update Tools"_
    4. Marca todas las casillas

4. Una vez acabado abre la **Paleta de Comandos** de nuevo
    1. Escribe: `shell`
    2. Elige: "Install 'code' command in PATH"
        1. **NOTA:** No tienes que hacer esto si estas en Windows.

## Eso es todo! Disfrutad! 🤩

<div style="page-break-after: always;"></div>

> Para mas tutoriales: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2021 Inanc Gumus
> 
> Learn Go Programming Course
> 
> [Click aqui para leer la licencia.](https://creativecommons.org/licenses/by-nc-sa/4.0/)