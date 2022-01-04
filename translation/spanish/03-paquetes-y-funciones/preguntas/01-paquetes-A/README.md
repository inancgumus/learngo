## ¿Dónde almacenar los archivos de código fuente que pertenecen a un paquete?
1. Cada archivo debe ir a un directorio diferente
2. En un solo directorio *CORRECTO*


## ¿Por qué se utiliza una cláusula de paquete en un archivo de código fuente de Go?
1. Se usa para importar un paquete.
2. Se utiliza para informarle a Go que el archivo pertenece a un paquete *CORRECTO*
3. Se usa para declarar una nueva función

> **1:** `import` declaración hace eso.
>
>
> **3:** `func` declaración hace eso.
>
>


## ¿Dónde debería poner el `package clause` en un archivo de código fuente de Go?
1. Como primer código en un archivo de código fuente de Go *CORRECTO*
2. Como último código en un archivo de código fuente de Go
3. Puedes ponerlo en cualquier lugar


## ¿Cuántas veces puede usar`package clause` para un solo archivo de código fuente?
1. Una vez *CORRECTO*
2. Ninguna
3. Varias veces


## ¿Cuál es el uso correcto de`package clause`?
1. `my package`
2. `package main`
3. `pkg main`


## ¿Cual es la correcta?
1. Todos los archivos pertenecen al mismo paquete no pueden llamar a las funciones de los demás
2. Todos los archivos pertenecen al mismo paquete pueden llamar a las funciones de los demás *CORRECTO*


## ¿Cómo ejecutar varios archivos Go?
1. go run *.*go
2. go build *go
3. go run go
4. go run *.go *CORRECTO*

> **4:** También puede llamarlo como (asumiendo que hay file1.go file2.go y file3.go en el mismo directorio): vaya a ejecutar file1.go file2.go file3.go
>
>
