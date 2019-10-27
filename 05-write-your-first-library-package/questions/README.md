# Questions: Library Packages #

---

## (1) Which of the following choices is true? ##

---

**Note**: Each answer choice has an explanation so even if you know the answer, select all of them one by one to read the explanations.

1. You can run a library packages.
2. In a library package, there should be a function named `main()`
3. You can compile library packages
4. You have to compile library packages.

---

**Solution: 3**

> 1. Library packages cannot be run; they can only be imported from other packages.
> 2. In library packages, you don't have to include the `main()` function because it isn't an executable package. Only executable packages need a `main()` function
> 4. Library packages don't need to be compiled. When you run your program, library packages will automatically be built by the library or package that it came fom.

---

## (2) Which of the following is true about exported names? ##

1. Exported names must be typed in all capital letters.
2. The first letter of an exported name must be a capital letter.
3. An exported name must be inside the function scope
4. A new file must be created just for the exported name.
---

**Solution: 2**

That's right. Only the first letter needs to be capitalized.

> 1. If a name is declared in all capital letters, it will be exported but it is not a **must**. It is only exported because the first letter is capitalized.
> 3. Exported names should be declared at the package scope, not the function scope.
> 4. Creating a new file just for the exported name is not necessary.

---

## (3) How can you access your library package's functions in an executable program? ##

1. The library package must be exported before you can access its imported names.
2. You need to import your library package first, then you can access it's exported names.
3. You can access your library package anywhere as long as it is in the same directory as your executable program
4. You can import it by using just by using the function's name.

---

**Solution: 2**

> 1. You can't access packages. All packages are already exported (_unless you put them in a directory called `"internal"`, but more on that later!_)
> 3. That's right. The package must be imported first before its functions are accessible.
> 4. You can't use functions from a library package just by calling the function's name. You need to import it by using the full directory path after the `$GOPATH`.

---


## (4) In the following program, which names are exported? ##

```go
package wizard

import "fmt"

func doMagic() {
    fmt.Println("enchanted!")
}

func Fireball() {
    fmt.Println("Fireball!!!")
}
```

1. `fmt`
2. `doMagic()`
3. Fireball()
4. `Println()`

---

**Solution: 3**

> 1. `fmt` is just the name of the imported package.
> 2. `doMagic()` does not start with a lowercase letter, so it is not exported.
> 3. Correct! It starts with an uppercase letter so it gets exported.
> 4. `Println()` is not exported because it has already been exported from `fmt` package.

---


## (5) In the following programs, which names are exported? ##

```go
package wizard
import "fmt"

var one string
var Two string
var greenTrees string

func doMagic() {
    fmt.Println("enchanted!")
}

func Fireball() {
    fmt.Println("fireball!!!")
}
```

1. `doMagic()` and `Fireball()`
2. `Fireball()` and `Two`
3. `Fireball()`, `greenTrees` and `Two`
4. `Fireball()`, `greenTrees`, `one` and `Two`

---

**Solution: 2**

That's right. `Fireball()` and `Two` both start with capital letters so the will be exported.

> 1. `doMagic()` starts with a lowercase letter, so it is not exported.
> 3. `greenTrees` starts with a lowercase letter as well.
> 4. `one` and `greenTrees` both do not start with capital letters, so they will not be exported.