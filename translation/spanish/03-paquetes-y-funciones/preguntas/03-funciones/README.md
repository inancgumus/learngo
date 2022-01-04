## ¿Qué es una funcion?
* Bloque de código ejecutable 
* La visibilidad de los nombres declarados **CORRECTO**
* Determina qué ejecutar

```go
package awesome

import "fmt"

var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```

## ¿Qué nombre a continuación tiene la fuincion del paquete(libreria)?
1. awesome
2. fmt
3. enabled **CORRECTO**
4. counter

> **3:** Eso es correcto.`enabled` enabledestá fuera de cualquier función, por lo que es un nombre de ámbito de paquete. La funcion `block()` también tiene el alcance del paquete; también está fuera de cualquier función.
>
>


## ¿Qué nombre a continuación tiene la funcion del archivo?
1. awesome
2. fmt **CORRECTO**
3. enabled
4. block()
5. counter

> **2:** Eso es correcto. Los nombres de los paquetes importados tienen un ámbito de archivo. Y solo se pueden usar dentro del mismo archivo.
>
>


## ¿Qué nombre a continuación está en el alcance de la función block ()?
1. awesome
2. fmt
3. enabled
4. block()
5. counter **CORRECTO*

> **5:** Eso es correcto.  `counter` se declara dentro de la funcion `block()` por lo que está en el alcance del bloque func. Fuera de la funcion `block()`, otro código no puede verlo.
>
>


## ¿Puede `block()` ver el nombre `enabled`?
1. Sí: está en el alcance del paquete **CORRECTO**
2. No: está en el alcance del archivo
3. No: está en el alcance del bloque de block ()

> **1:** Todo el código dentro del mismo paquete puede ver todos los demás nombres declarados a nivel de paquete.
>
>


## ¿Pueden otros archivos en el paquete `awesome` ver el nombre de `counter`?
1. Si
2. No: está en la funcion del paquete
3. No: está en la funcion del archivo
4. No: está en la funcion del block() **CORRECTO**

> **4:** Eso es correcto. Ninguno de los otros códigos puede ver los nombres dentro de la funcion `block()` .Solo el código dentro de la funcion `block()` puede verlos (Solo hasta cierto punto. Por ejemplo: Dentro del bloque, el código solo puede ver las variables declaradas antes).
>


## ¿Pueden otros archivos en el paquete `awesome` ver el nombre de`fmt` ?
1. Si
2. No: está en la funcion del paquete
3. No: está en la funcion del archivo **CORRECTO**
4. No: está en la funcion del bloque() 

> **3:** Solo el mismo archivo puede ver los paquetes importados, no los otros archivos, ya sea que estén en el mismo paquete o no.
>
>


## ¿Qué sucede si declaras el mismo nombre en el mismo ámbito que este?
```go
package awesome

import "fmt"

// declared twice in the package scope
var enabled bool
var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```
1. El nombre recién declarado prevalecerá sobre el anterior.
2. No puedo hacer eso. Ya se ha declarado en la funcion del paquete.  *CORRECTO*
3. No puedo hacer eso. Ya se ha declarado en la funcion del archivo.

> **2:** Eso es correcto. No puede declarar el mismo nombre en el mismo ámbito. Si pudiera hacerlo, ¿cómo volvería a acceder a él? ¿O a cuál?
>
>


## ¿Qué sucede si declaras el mismo nombre en otro ámbito como este?:
```go
package awesome

import "fmt"

// declared at the package scope
var enabled bool

func block() {
    // also declared in the block scope
    var enabled bool

    var counter int
    fmt.Println(counter)
}
```
1. El nombre recién declarado prevalecerá sobre el anterior. *CORRECTO*
2. No puedo hacer eso. Ya se ha declarado en la funcion del paquete.  
3. No puedo hacer eso. Ya se ha declarado en la funcion del archivo.

> **1:** En realidad, puede declarar el mismo nombre en los ámbitos internos de esta manera. La funcion `block()` está dentro de su paquete. Esto significa que puede acceder al alcance de su paquete (pero no al revés). Entonces, la funcion `block()`de 'está bajo el alcance de su paquete. Esto significa que puede volver a declarar el mismo nombre. Anulará el nombre del ámbito principal. Ambos pueden existir juntos. Consulte el ejemplo en el repositorio de cursos para averiguarlo.
>
>
