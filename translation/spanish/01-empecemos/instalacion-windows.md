# Instalación Windows

## NOTA

Si tienes el administrador de paquetes [chocolatey.org](https://chocolatey.org/) puedes instalar Go facilmente con este comando:

```bash
choco install golang
```

### 1. Instala El Editor Visual Studio Code

1. Instalalo pero no lo abras todavia
2. Ves a: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Elige la versión para Windows y empieza la descarga
4. Ejecuta el instalador

## 2. Instala Git

1. Descarga y ejecuta el instalador: Ves a: [https://gitforwindows.org](https://gitforwindows.org)
2. Elige VSCode como editor por defecto
3. Marca todas las casillas
4. Elige: "Utilizar Git desde la Consola de Comandos de Windows"
5. Codificaciones: Elige la opción: _"Dejar como esta..."

## 3. Instala Go

1. Ves a: [https://golang.org/dl](https://golang.org/dl)
2. Elige Windows y descargalo
3. Ejecuta el instalador

## 4. Configura Visual Studio Code

1. Abre VS Code; haz click en la pestaña de extensiones situada a la izquierda, busca "go" e instalalo
2. Cierra VS Code por completo y ábrelo de nuevo

3. Ve al menu Ver: selecciona **Paleta de Comandos**
    1. O simplemente presiona: `ctrl+shift+p`
    2. Escribe: `go install`
    3. Elige: _"Go: Install/Update Tools"_
    4. Marca todas las casillas

## 5. Usando Git-Bash

* En este curso hare uso de comandos bash. Bash es simplemente una linea de comandos usada en OS X y en Linux. Es una de las interfaces de linea de comandos mas populares. Asi que, si quieres usarla tambien, en vez de usar la linea de comandos por defecto de Windows, puedes usar Git-Bash. Con ella puedes escribir comandos tanto de OS X como de Linux.

* Si no quieres hacer uso de Git-Bash no pasa nada, depende de ti. Pero ten en cuenta que yo hare uso constante de comandos bash.

* Tambien puedes utilizar algo mas potente que Git-Bash como podria ser el: [Subsistema de Linux para Windows](https://docs.microsoft.com/es-es/windows/wsl/install-win10)

* **Para usar git bash sigue los siguientes pasos:**
    1. Busca git bash desde la barra de busqueda
    2. O, si hay un icono en tu escritorio, clicka en el

    3. Usa git bash por defecto en VS Code:
        1. Abre VS Code
        2. Ves a la **Paleta de Comandos**
            1. Escribe: `terminal`
            2. Elige: _"Terminal: Terminal por Defecto"_
            3. Y elige: _"Git Bash"_
        3. **NOTA:** Normalmente, puedes encontrar tus archivos en `c:\`, pero, cuando estas usando git bash los encontraras en `/c/`, que es el mismo directorio pero mas abreviada

## Eso es todo! Disfrutad! 🤩

<div style="page-break-after: always;"></div>

> Para mas tutoriales: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2021 Inanc Gumus
> 
> Learn Go Programming Course
> 
> [Click aqui para leer la licencia.](https://creativecommons.org/licenses/by-nc-sa/4.0/)